package repository

import (
	"database/sql"

	"github.com/AntonioFSRE/go-bid/internal/domain/models"
	"github.com/AntonioFSRE/go-bid/internal/domain/repositories"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type pgRepository struct {
	db *sqlx.DB
}

func NewPGRepository(db *sqlx.DB) repositories.PGBidRepository {
	return &pgRepository{db}
}

func (r *pgRepository) GetByID(id int64) (models.Bid, error) {
	var bid models.Bid

	if err := r.db.Get(
		&bid,
		getBidQuery,
		id,
	); err != nil {
		if err == sql.ErrNoRows {
			return bid, echo.ErrNotFound
		}

		return bid, echo.ErrBadRequest
	}

	return bid, nil
}

func (r *pgRepository) Store(a *models.Bid) (*models.Bid, error) {
	var bid models.Bid

	if err := r.db.QueryRowx(
		createBidQuery,
		a.ID,
		a.AuthorID,
		a.Ttl,
		a.Price,
		a.SetAt,
	).StructScan(&bid); err != nil {
		return nil, echo.ErrBadRequest
	}

	return &bid, nil
}

func (r *pgRepository) Update(a *models.Bid) (*models.Bid, error) {
	var bid models.Bid

	if err := r.db.QueryRowx(
		updateBidQuery,
		a.Price,
		a.AuthorID,
	).StructScan(&bid); err != nil {
		if err == sql.ErrNoRows {
			return nil, echo.ErrNotFound
		}

		return nil, echo.ErrBadRequest
	}

	return &bid, nil
}