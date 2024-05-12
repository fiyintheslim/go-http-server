package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Hello struct {
	Success bool `json:"success"`
}

func HandleHome(w http.ResponseWriter, r *http.Request) {

	res := Hello{
		Success: true,
	}
	val, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, err)
	}
	w.Header().Set("content-type", "application/json")
	w.Write(val)
}
