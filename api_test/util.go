package apitest

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"github.com/hotkimho/realworld-api/env"
	"github.com/hotkimho/realworld-api/types"
)

const TestHost = "http://localhost:8080"
const TestUsername = "kimhotest"
const TestEmail = "kimhotest@naver.com"
const TestPassword = "test"
const TestToken = ""
const TestTag = "test"
const TestAuthor = "1"
const TestArticleID = "1"
const TestTitle = "test title"
const TestDescription = "test description"
const TestBody = "test body"
const TestBio = "test bio"
const TestProfileImage = "test image"

// "localhost", "user", "{user_id}" 이런식으로 온걸 하나의 문자열로 합침
func makeURL(host string, params ...string) string {
	url := host
	for _, param := range params {
		url += "/" + param
	}

	return url
}

func RequestTest(method, url string, body io.Reader, header map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	return http.DefaultClient.Do(req)
}

func GetUserIDByJWT(token string) (string, error) {

	parsedToken, err := jwt.ParseWithClaims(token, &types.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.Config.Auth.Secret), nil
	})
	if err != nil {
		if strings.Compare(err.Error(), types.ERR_EXPIRED_TOKEN) == 0 {
			return "", errors.New("expired token")
		}
		return "", errors.New("invalid token")
	}

	if claim, ok := parsedToken.Claims.(*types.JWTClaims); ok && parsedToken.Valid {
		return fmt.Sprintf("%d", claim.UserID), nil

	}

	return "", errors.New("invalid token")
}
