package server

import (
	"fmt"
	"net/http"

	"github.com/fiyintheslim/go-http-server/internal/repository"
	"github.com/fiyintheslim/go-http-server/internal/routes"
)

func StartServer() {
	err := repository.InitDatabase()

	if err != nil {
		fmt.Println(err)
		return
	}
	defer repository.CloseDbConn()
	routes.DefineRoutes()
	http.ListenAndServe(":3030", nil)
}
