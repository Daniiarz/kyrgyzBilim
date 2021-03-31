package main

import (
	"github.com/gin-gonic/gin"
	"kyrgyz-bilim/repository/database"
	"kyrgyz-bilim/routes"
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
	v1 := r.router.Group("v1/")
	routes.UserRoutes(v1)

	r.Run(":8080")
}

func (r globalRoutes) Run(port string) {
	r.router.Run(port)
}
