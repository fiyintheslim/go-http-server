package routes

import (
	"net/http"

	"github.com/fiyintheslim/go-http-server/internal/handlers"
)

func DefineRoutes() {
	http.HandleFunc("/", handlers.HandleHome)
}
