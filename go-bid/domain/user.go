package domain

import "context"

// User ...
type User struct {
	userId    int64  `json:"id"`
	name      string `json:"name"`
	password  string `json:"password"`
	role      string `json:"role"`
}

// UserRepository represent the user's repository contract
type UserRepository interface {
	SignIn(ctx context.Context, username, password string) (string, error)
	CheckBid(ctx context.Context, userId int64) (User, error)
}

// UserUsecase represent the bid's usecases
type UserUsecase interface {
	SignIn(ctx context.Context, username, password string) (string, error)
}