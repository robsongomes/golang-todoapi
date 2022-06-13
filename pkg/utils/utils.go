package utils

import (
	"encoding/json"
	"net/http"
)

func ParseBody(req *http.Request, v interface{}) {
	json.NewDecoder(req.Body).Decode(&v)
}

func WriteJson(rw http.ResponseWriter, v interface{}) {
	json.NewEncoder(rw).Encode(v)
}
