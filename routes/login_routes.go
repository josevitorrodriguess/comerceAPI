package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/productsAPI/controller"
	"github.com/josevitorrodriguess/productsAPI/repository"
	"github.com/josevitorrodriguess/productsAPI/usecase"
)

func LoginRoutes(server *gin.Engine, dbConnection *sql.DB) {

	clientRepo := repository.NewClientRepository(dbConnection)
	merchantRepo := repository.NewMerchantRepository(dbConnection)

	
	clientAuthUsecase := usecase.NewAuthClientUsecase(clientRepo)
	merchantAuthUsecase := usecase.NewAuthMerchantUsecase(merchantRepo)

	authController := controller.NewAuthController(clientAuthUsecase, merchantAuthUsecase)

	
	server.POST("/login/client", authController.LoginClient)
	server.POST("/login/merchant", authController.LoginMerchant)
}
