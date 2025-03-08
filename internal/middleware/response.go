package middleware

import "net/http"

type DefaultResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewDefaultResponse(message string) *DefaultResponse {
	return &DefaultResponse{
		Status:  "UNAUTHORIZED",
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}
