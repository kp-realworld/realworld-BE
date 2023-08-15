package controller

import (
	"fmt"
	"net/http"
)

func TestFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
}
