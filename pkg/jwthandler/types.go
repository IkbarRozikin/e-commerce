package jwthandler

import (
	"product-service/internal/infrastructure"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

type CustomClaims struct {
	UserId     string `json:"user_id"`
	Role       string `json:"role"`
	IsVerified bool   `json:"is_verified"`
	jwt.RegisteredClaims
}

type CostumClaimsPayload struct {
	UserId          string    `json:"user_id"`
	Role            string    `json:"role"`
	IsVerified      bool      `json:"is_verified"`
	TokenExpiration time.Time `json:"token_expiration"`
}

func GenerateTokenString(payload CostumClaimsPayload) (string, error) {
	claims := CustomClaims{
		UserId:     payload.UserId,
		Role:       payload.Role,
		IsVerified: payload.IsVerified,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "user",
			Issuer:    "product-service",
			ExpiresAt: jwt.NewNumericDate(payload.TokenExpiration),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			NotBefore: jwt.NewNumericDate(time.Now().UTC()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, err := token.SignedString([]byte(infrastructure.Envs.Guard.JwtPrivateKey))
	if err != nil {
		log.Error().Err(err).Msg("jwthandler::GenerateTokenString - Error while signing token")
		return "", err
	}

	return tokenString, nil
}

func ParseTokenString(tokenString string) (*CustomClaims, error) {
	claims := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(infrastructure.Envs.Guard.JwtPrivateKey), nil
	})
	if err != nil {
		log.Error().Err(err).Msg("jwthandler::ParseTokenString - Error while parsing token")
		return nil, err
	}

	if !token.Valid {
		log.Error().Msg("jwthandler::ParseTokenString - Invalid token")
		return nil, err
	}

	return claims, nil
}
