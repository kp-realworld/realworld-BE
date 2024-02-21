package api_test

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	articledto "github.com/hotkimho/realworld-api/controller/dto/article"
	"io"
	"net/http"
	"testing"
)

func TestArticle(t *testing.T) {

	resp, err := http.Get("http://localhost:8080/articles")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code error: %d", resp.StatusCode)
	}

	// resp.body를 json으로 변환
	// json으로 변환한 body를 ArticleResponseDTO로 변환
	// 변환한 ArticleResponseDTO를 검증
	// 검증한 결과를 통해 테스트 결과를 판단
	// 테스트 결과에 따라 t.Fatalf로 에러를 출력
	// 테스트 결과에 따라 t.Logf로 로그를 출력

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	validate := validator.New()
	var articleResponseDTO articledto.ReadArticleByOffsetResponseWrapperDTO

	err = json.NewDecoder(resp.Body).Decode(&articleResponseDTO)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(articleResponseDTO)

	err = validate.Struct(articledto.ReadArticlesByUserResponseWrapperDTO{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("성공")
}
