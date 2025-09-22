package auth

import (
	"clean-arc-2/internal/infrastructure/logger"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		logger.GetLogger().Info(err.Error())
		return false, fmt.Errorf("Ошибка авторизации %w", err)
	}
	return true, nil
}

func ExtractPayloadFromToken(requestToken string, secret string) (*JwtCustomClaims, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		logger.GetLogger().Info(err.Error())
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Invalid Token")
	}

	jsonPayload, err := json.Marshal(claims)
	if err != nil {
		logger.GetLogger().Info(err.Error())
		return nil, fmt.Errorf("Error marshalling claims")
	}

	payload := &JwtCustomClaims{}
	if err = json.Unmarshal(jsonPayload, &payload); err != nil {
		logger.GetLogger().Info(err.Error())
		return nil, fmt.Errorf("Error unmarshalling claims")
	}

	return payload, nil
}
