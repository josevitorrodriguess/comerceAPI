package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/productsAPI/controller/middlewares"
	"github.com/josevitorrodriguess/productsAPI/model"
	"github.com/josevitorrodriguess/productsAPI/services"
	"github.com/josevitorrodriguess/productsAPI/usecase"
)

type productController struct {
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
   
	products, err := p.productUseCase.GetProducts()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to retrieve products",
            "details": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	// Executar o middleware de autenticação
	middlewares.Auth()(ctx)

	// Decodificar o JSON da requisição para a estrutura do produto
	var product model.Product
	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error decoding JSON"})
		return
	}

	// Obter as reivindicações (claims) do token do contexto
	tokenString := ctx.GetHeader("Authorization")[len("Bearer "):]
	claims, err := services.NewJWTService().GetClaims(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid JWT"})
		return
	}

	// Associar o ID do comerciante (MerchantID) ao produto com base nas reivindicações (claims) do token JWT
	product.MerchantID = claims.Sum

	// Chamar o caso de uso para criar o produto
	insertedProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error creating product"})
		return
	}

	// Retornar o produto criado com sucesso
	ctx.JSON(http.StatusCreated, insertedProduct)
}


// Método GetProductById do controller para recuperar um produto pelo ID
func (p *productController) GetProductById(ctx *gin.Context) {

	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "Product ID cannote be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "product Id must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "product was not found in the database",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) DeleteProduct(ctx *gin.Context) {

	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "Product ID cannote be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "product Id must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = p.productUseCase.DeleteProduct(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	response := model.Response{
		Message: "Product successfully deleted",
	}

	ctx.JSON(http.StatusOK, response)
}
