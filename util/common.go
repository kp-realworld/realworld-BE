package util

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetIntegerParam[T int | int16 | int32 | int64](r *http.Request, key string) (T, error) {
	vars := mux.Vars(r)
	val, ok := vars[key]
	if !ok {
		return 0, errors.New("invalid key")
	}

	parsedVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, err
	}

	var result T
	switch any(result).(type) {
	case int:
		result = T(parsedVal)
	case int16:
		result = T(int16(parsedVal))
	case int32:
		result = T(int32(parsedVal))
	case int64:
		result = T(int64(parsedVal))
	default:
		return 0, errors.New("invalid type")
	}

	return result, nil
}
