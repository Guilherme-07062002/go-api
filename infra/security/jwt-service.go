package security

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenService interface {
	GenerateToken(userId string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type JwtService struct {
	secretKey []byte
}

func NewJwtService(secret string) TokenService {
	return &JwtService{secretKey: []byte(secret)}
}

func (jwtService *JwtService) GenerateToken(userId string) (string, error) {
	// 1. Definimos os Claims (Payload do token)
	claims := jwt.RegisteredClaims{
		Subject:   userId,                                             // O "dono" do token
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Expira em 24h
		IssuedAt:  jwt.NewNumericDate(time.Now()),                     // Data de emissão
		Issuer:    "go-api",                                           // Quem emitiu o token
	}

	// 2. Criamos o token com o método de assinatura (HS256 é o padrão de mercado)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 3. Assinamos o token com a nossa chave secreta
	tokenString, err := token.SignedString(jwtService.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (jwtService *JwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtService.secretKey, nil
	})
}
