package http_utils

import (
	"encoding/json"
	"net/http"

	"github.com/amitabhprasad/bookstore-util-go/rest_errors"
)

func ResponseJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func ResponseError(w http.ResponseWriter, err rest_errors.RestErr) {
	ResponseJson(w, err.Status(), err)
}
