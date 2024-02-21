package apitest

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/go-playground/validator/v10"

	authdto "github.com/hotkimho/realworld-api/controller/dto/auth"
)

func TestSignUp(t *testing.T) {
	url := makeURL(TestHost, "user", "signup")
	// user

	signupReq := authdto.SignUpRequestDTO{
		Email:    TestEmail,
		Username: TestUsername,
		Password: TestPassword,
	}

	jsonByte, err := json.Marshal(signupReq)
	if err != nil {
		t.Fatal(err)
	}
	jsonBuf := bytes.NewBuffer(jsonByte)
	resp, err := RequestTest("POST", url, jsonBuf, nil)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		t.Fatalf("status code error: %d", resp.StatusCode)
	}

	dataBuf, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var signupRes authdto.SignUpResponseWrapperDTO
	err = json.Unmarshal(dataBuf, &signupRes)
	if err != nil {
		t.Fatal(err)
	}

	validate := validator.New()
	err = validate.Struct(signupRes)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSignin(t *testing.T) {
	url := makeURL(TestHost, "user", "signin")

	signinReq := authdto.SignInRequestDTO{
		Email:    TestEmail,
		Password: TestPassword,
	}

	jsonByte, err := json.Marshal(signinReq)
	if err != nil {
		t.Fatal(err)
	}

	jsonBuf := bytes.NewBuffer(jsonByte)
	resp, err := RequestTest("POST", url, jsonBuf, nil)
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

	var signinRes authdto.SignInResponseWrapperDTO
	err = json.Unmarshal(dataBuf, &signinRes)
	if err != nil {
		t.Fatal(err)
	}

	validate := validator.New()
	err = validate.Struct(signinRes)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerifyEmail(t *testing.T) {
	url := makeURL(TestHost, "user", "verify-email")

	verifyEmailReq := authdto.VerifyEmailRequestDTO{
		Email: TestEmail,
	}

	jsonByte, err := json.Marshal(verifyEmailReq)
	if err != nil {
		t.Fatal(err)
	}

	jsonBuf := bytes.NewBuffer(jsonByte)
	resp, err := RequestTest("POST", url, jsonBuf, nil)
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

	var verifyEmailRes authdto.VerifyEmailResponseDTO
	err = json.Unmarshal(dataBuf, &verifyEmailRes)
	if err != nil {
		t.Fatal(err)
	}

	validate := validator.New()
	err = validate.Struct(verifyEmailRes)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerifyUsername(t *testing.T) {
	url := makeURL(TestHost, "user", "verify-username")

	verifyUsernameReq := authdto.VerifyUsernameRequestDTO{
		Username: TestUsername,
	}

	jsonByte, err := json.Marshal(verifyUsernameReq)
	if err != nil {
		t.Fatal(err)
	}

	jsonBuf := bytes.NewBuffer(jsonByte)
	resp, err := RequestTest("POST", url, jsonBuf, nil)
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

	var verifyUsernameRes authdto.VerifyUsernameResponseDTO
	err = json.Unmarshal(dataBuf, &verifyUsernameRes)
	if err != nil {
		t.Fatal(err)
	}

	validate := validator.New()
	err = validate.Struct(verifyUsernameRes)
	if err != nil {
		t.Fatal(err)
	}
}
