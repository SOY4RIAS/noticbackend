package response

import (
	"encoding/json"
	"net/http"
)

func AsErrorJson(w http.ResponseWriter, code int, msg string) {
	AsJson(w, code, map[string]string{"message": msg})
}

func AsJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}
