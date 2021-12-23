package mongo

import (
	"context"
	"database/sql"

	"github.com/AntonioFSRE/go-bid/domain"
	"github.com/AntonioFSRE/go-bid/auth/repository"
	"github.com/golang-jwt/jwt"
)	


type UserRepository struct {
	 DB *sql.DB
}

func NewPostgreesUserRepository(db *sql.DB) domain.UserRepository {
	return &postgresUserRepo{
		DB: db,
	}
}

func (m *postgresUserRepo) GetUser(ctx context.Context, query string, args ...interface{}) (res domain.User, err error) {
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return domain.User{}, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	res = domain.User{}

	err = row.Scan(
		&res.userId,
		&res.username,
		&res.password,
		&res.role,
	)
	return
}

func toPostgresUser(u *domain.User) *User {
	return &User{
		Username: u.Username,
		Password: u.Password,
		Role: u.Role,
	}
}
