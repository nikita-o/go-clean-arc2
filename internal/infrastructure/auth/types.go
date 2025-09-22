package auth

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	ID        int    `json:"id"`
	SessionId string `json:"session_id"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	ID        int    `json:"id"`
	SessionId string `json:"session_id"`
	jwt.RegisteredClaims
}
