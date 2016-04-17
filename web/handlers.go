package web

import (
	"encoding/json"
	"net/http"
)

func status(w http.ResponseWriter, r *http.Request) {
	var returnMap map[string]interface{}
	encoder := json.NewEncoder(w)
	if encoder == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	returnMap["status"] = audioManager.Status()
	encoder.Encode(returnMap)
}
