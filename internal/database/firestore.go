package database

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type FirestoreDB struct {
	*firestore.Client
}

func (f *FirestoreDB) Connect() error {
	_, path, _, ok := runtime.Caller(0)

	if !ok {
		return fmt.Errorf("failed to read function caller")
	}

	credentials_path := filepath.Join(filepath.Dir(path), "..", "..", "firestore_credentials.json")
	opt := option.WithCredentialsFile(credentials_path)
	client, err := firestore.NewClient(context.Background(), os.Getenv("PROJECT_ID"), opt)

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
