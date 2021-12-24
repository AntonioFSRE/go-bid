package postgres

import (
	"context"
	"database/sql"

	"github.com/AntonioFSRE/go-bid/domain"
)

type postgresUserRepo struct {
	DB *sql.DB
}

// NewPostrgesUserRepository will create an implementation of user.Repository
func NewPostgresUserRepository(db *sql.DB) domain.UserRepository {
	return &postgresUserRepo{
		DB: db,
	}
}

func (m *postgresUserRepo) getOne(ctx context.Context, query string, args ...interface{}) (res domain.User, err error) {
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return domain.User{}, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	res = domain.User{}

	err = row.Scan(
		&res.userId,
		&res.username,
		&res.pasword,
		&res.role,
	)
	return
}

func (m *postgresUserRepo) GetByID(ctx context.Context, id int64) (domain.User, error) {
	query := `SELECT id, name, role FROM user WHERE userId=?`
	return m.getOne(ctx, query, userId)
}