package routes

import (
	"net/http"

	"github.com/fiyintheslim/go-http-server/internal/handlers"
)

func DefineRoutes() {
	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/user", handlers.HandleUser)
}
