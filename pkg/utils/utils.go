package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

// JSONResponse writes a JSON response
func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

// ErrorResponse writes an error response
func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// Extract Task ID from the URL
func ExtractTaskID(path string) (int, error) {
	parts := strings.Split(path, "/")
	if len(parts) > 3 {
		return strconv.Atoi(parts[3])
	}
	return 0, errors.New("invalid task ID")
}
