package utils

import (
	"encoding/json"
	"net/http"
)

const errorStatus = "ERROR"

type errorResponse struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

func FormatError(w http.ResponseWriter, statusCode int, err error) {
	SetJsonHeader(w)
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(errorResponse{
		Status: errorStatus,
		Reason: err.Error(),
	})
}
