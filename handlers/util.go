package handlers

import (
	"encoding/json"
	"github.com/iquark/go-api-sample/lib"
	"net/http"
)

// JSONResponse sends a response to the client formatting the incoming data to JSON
func JSONResponse(data []lib.Person, w http.ResponseWriter) {
	bytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}
