package utils

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON responds the http request
func RespondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
    response, err := json.MarshalIndent(payload, "", "  ")
    println("response is: ", response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}