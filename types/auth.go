package types

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

const (
	// JWT 토큰 만료 에러 문자열
	ERR_EXPIRED_TOKEN = "token has invalid claims: token is expired"
)
