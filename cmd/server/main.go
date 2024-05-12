package main

import (
	"github.com/fiyintheslim/go-http-server/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	
	server.StartServer()
}
