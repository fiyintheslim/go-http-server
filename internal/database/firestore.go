package database

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
)

type FirestoreDB struct {
	*firestore.Client
}

func (f *FirestoreDB) Connect() error {
	ctx := context.Background()
	client, err := firestore.NewClientWithDatabase(ctx, os.Getenv("PROJECT_ID"), os.Getenv("DATABASE_ID"))

	if err != nil {
		return fmt.Errorf("failed to initialize firestore database %v", err)
	}

	f.Client = client
	
	return nil
}

func (f * FirestoreDB) Close () {
	f.Client.Close()
}

func NewDBConnection() *FirestoreDB {
	return &FirestoreDB{}
}
