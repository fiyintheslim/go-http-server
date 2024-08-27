package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fiyintheslim/go-http-server/internal/entity"
	"github.com/fiyintheslim/go-http-server/internal/repository"
)

func HandleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handleCreateUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Internal server error marshal", http.StatusInternalServerError)
		return
	}

	var jsonData entity.CreateUserAccount

	uerr := json.Unmarshal(data, &jsonData)
	if uerr != nil {
		http.Error(w, "Internal server error marshal", http.StatusInternalServerError)
		return
	}

	fmt.Printf("body >>> %v", jsonData)
	user_repo := repository.NewUserRepository()

	user_repo.CreateUserAccount(&jsonData)
	
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
