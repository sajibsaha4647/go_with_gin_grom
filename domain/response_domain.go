package domain

import (
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    any    `json:"data"`
}

// Success sends a standard 200/201 structured response
func Success(message string, data any) Response {
	return Response{
		Message: message,
		Status:  http.StatusOK,
		Data:    data,
	}
}

// Error sends a standard error wrapper response
func Error(message string, statusCode int) Response {
	return Response{
		Message: message,
		Status:  statusCode,
		Data:    nil,
	}
}