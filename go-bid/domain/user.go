package domain

import "context"

// User ...
type User struct {
	UserId    int64  `json:"userId"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

// UserRepository represent the user's repository contract
type UserRepository interface {
	GetByID(ctx context.Context, UserId int64) (User, error)
	
}

// UserUsecase represent the bid's usecases
type UserUsecase interface {
	SignIn(ctx context.Context, username, password string) (string, error)
}