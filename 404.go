package main

import (
	"net/http"
)

func (opt Command) NotFound(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"description": "Unknown request",
		"error_code":  "CF-NotFound",
		"code":        10000,
	}

	http.Error(w, ConvertToJSON(response), http.StatusNotFound)
}
