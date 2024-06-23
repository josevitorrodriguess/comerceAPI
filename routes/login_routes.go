package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/productsAPI/controller"
	"github.com/josevitorrodriguess/productsAPI/repository"
	"github.com/josevitorrodriguess/productsAPI/usecase"
)

func LoginRoutes(server *gin.Engine, dbConnection *sql.DB) {
	MerchantRepository := repository.NewMerchantRepository(dbConnection)
	authUsecaseMerchant := usecase.NewAuthUsecase(MerchantRepository)
	authController := controller.NewAuthController(authUsecaseMerchant)


	server.POST("/login", authController.Login)

}
