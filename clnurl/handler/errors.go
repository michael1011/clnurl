package handler

import (
	"encoding/json"
	"net/http"
)

const errorStatus = "ERROR"

type errorResponse struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

func formatError(w http.ResponseWriter, statusCode int, err error) {
	setHeaders(w)
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(errorResponse{
		Status: errorStatus,
		Reason: err.Error(),
	})
}
