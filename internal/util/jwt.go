package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secret             string
	expiryHours        int
	refreshExpiryHours int
}

type TokenClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func NewJWTManager(secret string, expiryHours, refreshExpiryHours int) *JWTManager {
	return &JWTManager{
		secret:             secret,
		expiryHours:        expiryHours,
		refreshExpiryHours: refreshExpiryHours,
	}
}

func (jm *JWTManager) GenerateToken(userID, email string) (string, error) {
	claims := TokenClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jm.expiryHours))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "spotify-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jm.secret))
}

func (jm *JWTManager) GenerateRefreshToken(userID, email string) (string, error) {
	claims := TokenClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jm.refreshExpiryHours))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "spotify-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jm.secret))
}

func (jm *JWTManager) VerifyToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(jm.secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
