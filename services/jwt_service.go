package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwtService encapsula a lógica para geração e validação de tokens JWT
type jwtService struct {
	secretKey string // Chave secreta para assinar e validar tokens JWT
	issuer    string // Emissor dos tokens JWT
}

// NewJWTService cria e retorna uma nova instância do serviço JWT com a chave secreta configurada
func NewJWTService() *jwtService {
	return &jwtService{
		secretKey: os.Getenv("JWT_SECRET_KEY"), // Carrega a chave secreta do ambiente
		issuer:    "commerce-api-merchant",    // Define o emissor dos tokens JWT
	}
}

// Claim define as reivindicações personalizadas que serão incluídas nos tokens JWT
type Claim struct {
	Sum int `json:"sum"` // Dado personalizado na reivindicação
	jwt.StandardClaims   // Reivindicações padrão (expiração, emissão, etc.)
}

// GenerateToken gera um novo token JWT com base no ID fornecido
func (s *jwtService) GenerateToken(id int) (string, error) {
	// Criação das reivindicações personalizadas (claims)
	claim := &Claim{
		Sum: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), 
			Issuer:    s.issuer,                             
			IssuedAt:  time.Now().Unix(),                    
		},
	}

	// Cria um novo token JWT com as reivindicações definidas e o método de assinatura HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// Assina o token com a chave secreta configurada no serviço JWT
	tokenString, err := token.SignedString([]byte(s.secretKey))
	return tokenString, err
}

// ValidateToken verifica se o token JWT fornecido é válido
func (s *jwtService) ValidateToken(tokenString string) bool {
	// Faz o parse do token JWT fornecido
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifica se o método de assinatura do token é HMAC (HS256 neste caso)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token")
		}
		// Retorna a chave secreta para validar o token
		return []byte(s.secretKey), nil
	})

	
	return err == nil && token.Valid
}
