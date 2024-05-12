package repository

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/fiyintheslim/go-http-server/internal/database"
)

var (
	dBConn *firestore.Client

	dbCTX context.Context
)

func InitDatabase() error {
	fdb := database.NewDBConnection()
	if err := fdb.Connect(); err != nil {
		return err
	}
	fmt.Println("Database connection established")
	dBConn = fdb.Client
	dbCTX = context.Background()

	return nil
}

func CloseDbConn() {
	if dBConn != nil {
		fmt.Println("Database connection closed")
		dBConn.Close()
	}
}
