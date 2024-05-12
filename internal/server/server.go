package server

import (
	"net/http"

	"github.com/fiyintheslim/go-http-server/internal/repository"
	"github.com/fiyintheslim/go-http-server/internal/routes"
)

func StartServer() error {
	err := repository.InitDatabase()

	if err != nil {
		return err
	}
	defer repository.CloseDbConn()
	routes.DefineRoutes()
	http.ListenAndServe(":3030", nil)
	return nil
}
