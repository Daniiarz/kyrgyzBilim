package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"kyrgyz-bilim/repository/database"
	"kyrgyz-bilim/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	v1 := r.router.Group("/v1")
	routes.AuthRouters(v1)
	routes.UserRouters(v1)
	routes.CourseRoutes(v1)
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
