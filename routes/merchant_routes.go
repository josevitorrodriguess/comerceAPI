package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/productsAPI/controller"
	"github.com/josevitorrodriguess/productsAPI/repository"
	"github.com/josevitorrodriguess/productsAPI/usecase"
)

func MerchantRoutes(server *gin.Engine, dbConnection *sql.DB) {

	MerchantRepository := repository.NewMerchantRepository(dbConnection)
	MerchantUseCase := usecase.NewMerchantUsecase(MerchantRepository)
	MerchantController := controller.NewMerchantController(MerchantUseCase)

	server.GET("/merchants", MerchantController.GetMerchants)
	server.GET("/merchant/:merchantId", MerchantController.GetMerchantByID)
	server.POST("/merchant/", MerchantController.CreateMerchant)
	server.DELETE("/merchant/delete/:merchantId", MerchantController.DeleteMerchant)
}
