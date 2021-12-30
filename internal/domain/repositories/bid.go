package repositories

import (
	"time"

	"github.com/AntonioFSRE/go-bid/internal/domain/models"
)

type (
	PGBidRepository interface {
		GetByID(id int64) (models.Bid, error)
		Store(a *models.Bid) (*models.Bid, error)
		Update(a *models.Bid) (*models.Bid, error)
	}

	RedisBidRepository interface {
		GetByID(id int64) (models.Bid, error)
		SetBid(bid *models.Bid, exp time.Duration) error
	}
)
