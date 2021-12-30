package usecases

import (
	"github.com/AntonioFSRE/go-bid/internal/domain/models"
	"github.com/google/uuid"
)

type (
	jwtUseCase interface {
		Auth(user *models.User) (*models.AuthUser, error)
		GetToken(id uuid.UUID, tokenID uuid.UUID) (uuid.UUID, error)
	}

	UserUseCase interface {
		GetByID(id uuid.UUID) (models.User, error)
		Login(user *models.User) (*models.AuthUser, error)
		jwtUseCase
	}
)
