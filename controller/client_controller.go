package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/productsAPI/model"
	"github.com/josevitorrodriguess/productsAPI/usecase"
)

type clientController struct {
	clientUsecase usecase.ClientUsecase
}

func NewClientController(usecase usecase.ClientUsecase) clientController {
	return clientController{
		clientUsecase: usecase,
	}
}


func (c *clientController) GetClients(ctx *gin.Context) {
	clients, err := c.clientUsecase.GetClients()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, clients)
}


func (c *clientController) GetClientByID(ctx *gin.Context) {
	id := ctx.Param("clientId")
	if id == "" {
		response := model.Response{
			Message: "client ID cannot be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	clientId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "client ID must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	client, err := c.clientUsecase.GetClientByID(clientId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if client == nil {
		response := model.Response{
			Message: "client was not found in the database",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, client)
}


func (c *clientController) CreateClient(ctx *gin.Context) {
	var client model.Client
	err := ctx.BindJSON(&client)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertedClient, err := c.clientUsecase.CreateClient(client)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, insertedClient)
}

func (c *clientController) DeleteClient(ctx *gin.Context) {
	id := ctx.Param("clientId")
	if id == "" {
		response := model.Response{
			Message: "client ID cannot be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	clientId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "client ID must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = c.clientUsecase.DeleteClient(clientId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	response := model.Response{
		Message: "client successfully deleted",
	}

	ctx.JSON(http.StatusOK, response)
}
