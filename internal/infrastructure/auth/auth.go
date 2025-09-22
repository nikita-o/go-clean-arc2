package auth

import (
	"clean-arc-2/internal/infrastructure/logger"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func CreateAccessToken(userId int, sessionId string, secret string, expiry int) (accessToken string, err error) {
	now := time.Now().UTC()
	exp := now.Add(time.Minute * time.Duration(expiry))
	claims := &JwtCustomClaims{
		ID:        userId,
		SessionId: sessionId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		logger.GetLogger().Info(err.Error())
		return "", fmt.Errorf("Signed token error")
	}

	return t, err
}

func CreateRefreshToken(userId int, secret string, expiry int, sessionId string) (refreshToken string, err error) {
	exp := time.Now().UTC().Add(time.Hour * time.Duration(expiry))
	claimsRefresh := &JwtCustomRefreshClaims{
		ID:        userId,
		SessionId: sessionId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		logger.GetLogger().Info(err.Error())
		return "", fmt.Errorf("Signed token error")
	}
	return rt, err
}
