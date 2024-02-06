package util

import (
	"errors"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/hotkimho/realworld-api/env"
	"github.com/hotkimho/realworld-api/types"
)

// 로그인 시, jwt 토큰 발급 함수
func IssueJWT(userID int64) (string, error) {
	if userID <= 0 {
		return "", errors.New("user id is empty")
	}

	now := time.Now().In(env.Seoul)
	ExpiredTime := now.Add(time.Hour * time.Duration(env.Config.Auth.AccessTokenExpire))
	claims := types.JWTClaims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(ExpiredTime),
			Issuer:    env.Config.Auth.Issuer,
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(env.Config.Auth.Secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyJWT(token string) (*types.JWTClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &types.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.Config.Auth.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claim, ok := parsedToken.Claims.(*types.JWTClaims); ok && parsedToken.Valid {
		return claim, nil
	}

	return nil, errors.New("token is invalid")
}

func VerifyEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return regex.MatchString(email)
}
