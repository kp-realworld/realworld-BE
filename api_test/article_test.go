package apitest

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/go-playground/validator/v10"

	articledto "github.com/hotkimho/realworld-api/controller/dto/article"
)

func TestReadArticles(t *testing.T) {
	url := makeURL(TestHost, "articles")

	resp, err := RequestTest("GET", url, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code error: %d", resp.StatusCode)
	}

	dataBuf, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var readArticlesRes articledto.ReadArticleByOffsetResponseWrapperDTO

	err = json.Unmarshal(dataBuf, &readArticlesRes)
	if err != nil {
		t.Fatal(err)
	}

	validate := validator.New()
	err = validate.Struct(readArticlesRes)
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadArticlesByTag(t *testing.T) {
	url := makeURL(TestHost, "articles", "tag")

	header := make(map[string]string)
	header["tag"] = TestTag

	resp, err := RequestTest("GET", url, nil, header)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code error: %d", resp.StatusCode)
	}

	dataBuf, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var readArticlesRes articledto.ReadArticleByOffsetResponseWrapperDTO

	err = json.Unmarshal(dataBuf, &readArticlesRes)
	if err != nil {
		t.Fatal(err)
	}

	validate := validator.New()
	err = validate.Struct(readArticlesRes)
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadMyArticles(t *testing.T) {
	url := makeURL(TestHost, "my", "articles")

	header := make(map[string]string)
	header["Authorization"] = TestToken

	resp, err := RequestTest("GET", url, nil, header)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code error: %d", resp.StatusCode)
	}

	dataBuf, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var readArticlesRes articledto.ReadArticleByOffsetResponseWrapperDTO

	err = json.Unmarshal(dataBuf, &readArticlesRes)
	if err != nil {
		t.Fatal(err)
	}

	validate := validator.New()
	err = validate.Struct(readArticlesRes)
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadArticlesByUser(t *testing.T) {
	url := makeURL(TestHost, "user", TestAuthor, "articles")

	resp, err := RequestTest("GET", url, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code error: %d", resp.StatusCode)
	}

	dataBuf, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var readArticlesRes articledto.ReadArticleByOffsetResponseWrapperDTO

	err = json.Unmarshal(dataBuf, &readArticlesRes)
	if err != nil {
		t.Fatal(err)
	}

	validate := validator.New()
	err = validate.Struct(readArticlesRes)
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadArticleByAuthor(t *testing.T) {
	authorID, err := GetUserIDByJWT(TestToken)
	if err != nil {
		t.Fatal(err)
	}

	url := makeURL(TestHost, "user", authorID, "article", TestArticleID)

	resp, err := RequestTest("GET", url, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code error: %d", resp.StatusCode)
	}

	dataBuf, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var readArticleRes articledto.ReadArticleResponseWrapperDTO

	err = json.Unmarshal(dataBuf, &readArticleRes)
	if err != nil {
		t.Fatal(err)
	}

	validate := validator.New()
	err = validate.Struct(readArticleRes)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateArticle(t *testing.T) {
	userID, err := GetUserIDByJWT(TestToken)
	if err != nil {
		t.Fatal(err)
	}

	url := makeURL(TestHost, "user", userID, "article")

	createArticleReq := articledto.CreateArticleRequestDTO{
		Title:       TestTitle,
		Body:        TestBody,
		Description: TestDescription,
		TagList:     []string{TestTag},
	}

	jsonByte, err := json.Marshal(createArticleReq)
	if err != nil {
		t.Fatal(err)
	}

	header := make(map[string]string)
	header["Authorization"] = TestToken

	jsonBuf := bytes.NewBuffer(jsonByte)
	resp, err := RequestTest("POST", url, jsonBuf, header)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code error: %d", resp.StatusCode)
	}

	dataBuf, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var createArticleRes articledto.CreateArticleResponseWrapperDTO
	err = json.Unmarshal(dataBuf, &createArticleRes)
	if err != nil {
		t.Fatal(err)
	}

	validate := validator.New()
	err = validate.Struct(createArticleRes)
	if err != nil {
		t.Fatal(err)
	}
}
