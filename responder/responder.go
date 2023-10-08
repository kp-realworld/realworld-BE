package responder

import "net/http"

func Response(w http.ResponseWriter, statusCode int, message string) {

	w.WriteHeader(statusCode)
	w.Write([]byte(message))

}
