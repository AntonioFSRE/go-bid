package postgres

import (
	"context"
	"database/sql"

	"github.com/AntonioFSRE/go-bid/domain"
)	

type postgresUserRepository struct {
	Conn *sql.DB
}

// NewPostgresBidRepository will create an object that represent the bid.Repository interface
func NewPostgresUserRepository(Conn *sql.DB) domain.UserRepository {
	return &postgresUserRepository{Conn}
}

func (m *postgresUserRepository) GetByID(ctx context.Context, userId int64) (domain.User, error) {
	query := `SELECT userId, name, role FROM user WHERE userId=?`
	return m.GetUser(ctx, query, userId)
}

func (m *postgresUserRepository) GetUser(ctx context.Context, query string, args ...interface{}) (res domain.User, err error) {
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return domain.User{}, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	res = domain.User{}

	err = row.Scan(
		&res.UserId,
		&res.Name,
		&res.Password,
		&res.Role,
	)
	return toDomain(user), nil
}

func toPostgresUser(u *domain.User) *User {
	return &User{
		Username: u.Name,
		Password: u.Password,
		Role: u.Role,
	}
}	
func toDomain(u *User) *domain.User {
	return &domain.User{
		UserId: u.UserId.Hex(),
		Name: u.Name,
		Password: u.Password,
		Role: u.Role,
	}
}

