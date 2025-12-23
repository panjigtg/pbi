package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"pbi/internal/pkg/models"
)

var (
	jwtSecret []byte
	jwtExpire time.Duration
)

func InitJWT(secret string, expire time.Duration) {
	jwtSecret = []byte(secret)
	jwtExpire = expire
}

func GenerateToken(userID int, isAdmin bool) (string, error) {
	claims := models.JWTClaims{
		UserID:  userID,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*models.JWTClaims, error) {
	if len(jwtSecret) == 0 {
		return nil, errors.New("JWT not initialized")
	}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&models.JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return jwtSecret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
