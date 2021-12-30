package repository

import (
	"database/sql"
	"net/http"

	"github.com/AntonioFSRE/go-bid/internal/domain/models"
	"github.com/AntonioFSRE/go-bid/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type pgRepository struct {
	db *sqlx.DB
}

func NewPGRepository(db *sqlx.DB) repositories.PGUserRepository {
	return &pgRepository{db}
}

func (r *pgRepository) GetByID(id uuid.UUID) (models.User, error) {
	var user models.User

	if err := r.db.Get(
		&user,
		getUserQuery,
		id,
	); err != nil {
		if err == sql.ErrNoRows {
			return user, echo.ErrNotFound
		}

		return user, echo.ErrBadRequest
	}

	return user, nil
}
func (r *pgRepository) FindByName(name string) (models.User, error) {
	var user models.User

	if err := r.db.QueryRowx(
		findUserByNameQuery,
		name,
	).StructScan(&user); err != nil {
		if err == sql.ErrNoRows {
			return user, echo.NewHTTPError(
				http.StatusBadRequest,
				"user is not found",
			)
		}

		return user, echo.ErrBadRequest
	}

	return user, nil
}