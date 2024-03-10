package utils 

import (
	"encoding/json"
	"net/http"

	"github.com/TheAnswer16/rapper_phrases_generator/models"
)

func ErrorHandler(w http.ResponseWriter, message string, statusCode int, err error) {

	var errorResponse models.Error

	errorResponse.Message = message
	errorResponse.StatusCode = statusCode
	errorResponse.Error = err.Error() 

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorResponse)
}