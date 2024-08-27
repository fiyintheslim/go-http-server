package entity

import "time"

type UserAccount struct {
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
	DeletedAt int64 `json:"deletedAt"`

	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	PasswordHash  string `json:"password"`
}

type CreateUserAccount struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	PasswordHash  string `json:"password"`
}

func NewUserAccount() *UserAccount {
	return &UserAccount{
		CreatedAt: time.Now().UnixMilli(),
		UpdatedAt: time.Now().UnixMilli(),
		DeletedAt: 0,
	}
}
