package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/productsAPI/controller"
	"github.com/josevitorrodriguess/productsAPI/repository"
	"github.com/josevitorrodriguess/productsAPI/usecase"
)

func ClientRoutes(server *gin.Engine, dbConnection *sql.DB) {

	ClientRepository := repository.NewClientRepository(dbConnection)
	ClientUsecase := usecase.NewClientUsecase(ClientRepository)
	ClientController := controller.NewClientController(ClientUsecase)


	server.GET("/clients", ClientController.GetClients)
	server.GET("/client/:clientId", ClientController.GetClientByID)
	server.POST("/client", ClientController.CreateClient)
	server.DELETE("/client/delete/:clientId", ClientController.DeleteClient)
	

}
