package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/productsAPI/model"
	"github.com/josevitorrodriguess/productsAPI/usecase"
)

type AuthController struct {
	clientAuthUsecase   usecase.AuthClientUsecase
	merchantAuthUsecase usecase.AuthMerchantUseCase
}

func NewAuthController(clientAuthUsecase usecase.AuthClientUsecase, merchantAuthUsecase usecase.AuthMerchantUseCase) AuthController {
	return AuthController{
		clientAuthUsecase:   clientAuthUsecase,
		merchantAuthUsecase: merchantAuthUsecase,
	}
}

func (ac *AuthController) LoginClient(ctx *gin.Context) {
	var l model.Login
	if err := ctx.BindJSON(&l); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "cannot bind JSON: " + err.Error()})
		return
	}

	token, err := ac.clientAuthUsecase.LoginClient(l.Email, l.Password)
	if err != nil {
		if err.Error() == "cannot find merchant" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "cannot find client"})
		} else if err.Error() == "invalid credentials" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (ac *AuthController) LoginMerchant(ctx *gin.Context) {
	var l model.Login
	if err := ctx.BindJSON(&l); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "cannot bind JSON: " + err.Error()})
		return
	}

	token, err := ac.merchantAuthUsecase.Login(l.Email, l.Password)
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
