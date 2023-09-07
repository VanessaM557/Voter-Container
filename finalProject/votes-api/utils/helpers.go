package utils

import (
	"encoding/json"
	"net/http"
	"strings"
)

func RespondJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func RespondError(w http.ResponseWriter, statusCode int, message string) {
	RespondJSON(w, statusCode, map[string]string{"error": message})
}

func ParseIDFromURL(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}

func IsValidID(id string) bool {
	for _, char := range id {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
