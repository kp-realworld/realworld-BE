package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/hotkimho/realworld-api/domain"
	"github.com/hotkimho/realworld-api/responder"
	"net/http"
)

var Validate *validator.Validate

type InsertUserRequest struct {
	Username string `json:"username" validate:"required"`
	Age      int    `json:"age" validate:"required, gte=0"`
}

func TestFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
	var data InsertUserRequest

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("err", err)
		responder.Response(w, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("data", data)

	err = Validate.Struct(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := domain.Users.Insert()
	if err != nil {
		responder.Response(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("id", id)

	responder.Response(w, http.StatusOK, "success")
}
