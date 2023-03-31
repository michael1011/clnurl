package utils

import "net/http"

func SetJsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
