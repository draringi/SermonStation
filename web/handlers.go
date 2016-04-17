package web

import (
	"encoding/json"
	"net/http"
)

func status(w http.ResponseWriter, r *Request) {
	var returnMap map[string]interface{}
	encoder := json.NewEncoder(w)
	if encoder == nil {
		w.SetHeader(StatusInternalServerError)
		return
	}
	returnMap["status"] = manager.Status()
	encoder.Encode(returnMap)
}
