package usecases

import (
	"github.com/AntonioFSRE/go-bid/internal/domain/models"
)

type BidUseCase interface {
	GetByID(id int64) (models.Bid, error)
	Store(a *models.Bid) (*models.Bid, error)
	Update(a *models.Bid) (*models.Bid, error)
}
