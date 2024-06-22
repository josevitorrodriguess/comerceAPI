package main

import (
	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/productsAPI/controller"
	"github.com/josevitorrodriguess/productsAPI/db"
	"github.com/josevitorrodriguess/productsAPI/repository"
	"github.com/josevitorrodriguess/productsAPI/usecase"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository :=  repository.NewProductRepository(dbConnection)

	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	
	ProductController := controller.NewProductControlller(ProductUseCase)
	
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product",ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)

	server.Run(":8000")
}