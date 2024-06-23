package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/productsAPI/services"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema = "Bearer "

		// Obter o cabeçalho Authorization da requisição
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing JWT"})
			ctx.Abort()
			return
		}

		// Extrair o token JWT removendo o prefixo "Bearer "
		token := header[len(BearerSchema):]

		// Validar o token JWT usando o serviço JWT
		jwtService := services.NewJWTService()
		if !jwtService.ValidateToken(token) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid JWT"})
			ctx.Abort()
			return
		}

		// Obter as reivindicações (claims) do token e adicioná-las ao contexto
		claims, err := jwtService.GetClaims(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid JWT claims"})
			ctx.Abort()
			return
		}
		ctx.Set("claims", claims)

		// Continuar para o próximo middleware ou manipulador se o token for válido
		ctx.Next()
	}
}
