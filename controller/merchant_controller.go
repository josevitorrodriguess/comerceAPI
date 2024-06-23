package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/productsAPI/model"
	"github.com/josevitorrodriguess/productsAPI/usecase"
)

type merchantController struct {
	merchantUsecase usecase.MerchantUsecase
}

func NewMerchantController(usecase usecase.MerchantUsecase) merchantController {
	return merchantController{
		merchantUsecase: usecase,
	}
}

func (m *merchantController) GetMerchants(ctx *gin.Context) {

	merchants, err := m.merchantUsecase.GetMerchants()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, merchants)
}

func (m *merchantController) GetMerchantByID(ctx *gin.Context) {

	id := ctx.Param("merchantId")
	if id == "" {
		response := model.Response{
			Message: "merchant ID cannote be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
	}

	merchantId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "merchant Id must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	merchant, err := m.merchantUsecase.GetMerchantById(merchantId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if merchant == nil {
		response := model.Response{
			Message: "merchant was not found in the database",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, merchant)
}

func (m *merchantController) CreateMerchant(ctx *gin.Context) {

	var merchant model.Merchant
	err := ctx.BindJSON(&merchant)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertMerchant, err := m.merchantUsecase.CreateMerchant(merchant)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertMerchant)
}

func (m *merchantController) DeleteMerchant(ctx *gin.Context) {

	id := ctx.Param("merchantId")
	if id == "" {
		response := model.Response{
			Message: "merchant ID cannot be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	merchantId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "merchant Id must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = m.merchantUsecase.DeleteMerchant(merchantId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	response := model.Response{
		Message: "merchant successfully deleted",
	}

	ctx.JSON(http.StatusOK, response)
}
