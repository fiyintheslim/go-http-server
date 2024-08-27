package repository

import (
	"fmt"

	"github.com/fiyintheslim/go-http-server/internal/entity"
)

type UserAccountRepository struct{}

func (r *UserAccountRepository) CollectionName() string {
	return "users-collection"
}

func (r *UserAccountRepository) CreateUserAccount(d *entity.CreateUserAccount) error {
	u := entity.NewUserAccount()
	u.FirstName = d.FirstName
	u.LastName = d.LastName
	u.Email = d.Email
	u.PasswordHash = d.PasswordHash

	docRef := dBConn.Collection(r.CollectionName()).NewDoc()

	u.ID = docRef.ID

	_, err := docRef.Set(dbCTX, u)

	if err != nil {
		return fmt.Errorf("failed to create a new user: %v", err)
	}

	return nil
}

func NewUserRepository() *UserAccountRepository {
	return &UserAccountRepository{}
}
