package main

import (
	"context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/gin-gonic/gin"
	"kyrgyz-bilim/repository/database"
	"kyrgyz-bilim/routes"
	"kyrgyz-bilim/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"                 // Import the adapter, it must be imported. If it is not imported, you need to define it yourself.
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres" // Import the sql driver
	_ "github.com/GoAdminGroup/themes/adminlte"                      // Import the theme
)

type globalRoutes struct {
	router *gin.Engine
}

func main() {
	database.DB = database.Connect()
	database.SetupDB(database.DB)
	r := globalRoutes{
		router: gin.Default(),
	}
	r.router.MaxMultipartMemory = 8 << 20
	v1 := r.router.Group("v1/")
	routes.AuthRouters(v1)
	routes.UserRouters(v1)
	routes.CourseRoutes(v1)

	adminCfg := config.Config{
		Databases: config.DatabaseList{
			"default": config.Database{
				Host:       os.Getenv("POSTGRES_HOST"),
				Port:       "5432",
				User:       os.Getenv("POSTGRES_USER"),
				Pwd:        os.Getenv("POSTGRES_PASSWORD"),
				Name:       os.Getenv("POSTGRES_DB"),
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     "postgresql",
			},
		},
		UrlPrefix: "admin", // The url prefix of the website.
		// Store must be set and guaranteed to have write access, otherwise new administrator users cannot be added.
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language: language.EN,
	}

	eng := engine.Default()
	_ = eng.AddConfig(adminCfg).
		AddGenerators(datamodel.Generators).
		AddGenerator("users", utils.GetUsersTable).
		AddGenerator("sections", utils.GetSectionsTable).
		AddGenerator("topics", utils.GetTopicsTable).
		AddGenerator("sub-topics", utils.GetSubTopicsTable).
		AddGenerator("courses", utils.GetCoursesTable).
		Use(r)
	r.Run(":8080")
}

func (r globalRoutes) Run(port string) {
	srv := &http.Server{
		Addr:    port,
		Handler: r.router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 3 seconds.")
	}
	log.Println("Server exiting")

}
