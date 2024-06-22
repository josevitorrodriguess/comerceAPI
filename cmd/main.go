package main

import (
	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/productsAPI/db"
	"github.com/josevitorrodriguess/productsAPI/routes"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	routes.ProductRoutes(server, dbConnection)

	server.Run(":8000")
}
