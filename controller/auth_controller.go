package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/productsAPI/model"
	"github.com/josevitorrodriguess/productsAPI/usecase"
)

type AuthController struct {
	authUsecase usecase.AuthUSeCase
}


func NewAuthController(authUsecase usecase.AuthUSeCase) AuthController {
	return AuthController{authUsecase}
}

func (ac *AuthController) Login(ctx *gin.Context) {
	var l model.Login
	if err := ctx.BindJSON(&l); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "cannot bind JSON" + err.Error()})
	}

	token, err := ac.authUsecase.Login(l.Email, l.Password)
	if err != nil {
		if err.Error() == "cannot find merchant" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else if err.Error() == "invalid credentials" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
