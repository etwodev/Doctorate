package helpers

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message}, "application/json")
}

func RespondWithCode(w http.ResponseWriter, httpcode int, code string) {
	RespondWithJSON(w, httpcode, map[string]string{"result": code}, "application/json")
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}, ct string) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", ct)
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithRaw(w http.ResponseWriter, code int, payload []byte, ct string) {
	w.Header().Set("Content-Type", ct)
	w.WriteHeader(code)
	w.Write(payload)
}