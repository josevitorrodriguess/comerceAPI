package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/productsAPI/controller"
	"github.com/josevitorrodriguess/productsAPI/controller/middlewares"
	"github.com/josevitorrodriguess/productsAPI/repository"
	"github.com/josevitorrodriguess/productsAPI/usecase"
)

func ProductRoutes(server *gin.Engine, dbConnection *sql.DB) {
	productRepository := repository.NewProductRepository(dbConnection)
	productUseCase := usecase.NewProductUseCase(productRepository)
	productController := controller.NewProductController(productUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)
	server.GET("/product/:productId", productController.GetProductById)
	server.DELETE("/product/delete/:productId", productController.DeleteProduct)

	// Grupo de rotas que necessitam de autenticação
	authProductRoutes := server.Group("/products", middlewares.Auth())
	{
		authProductRoutes.POST("/", productController.CreateProduct)
	}
}
