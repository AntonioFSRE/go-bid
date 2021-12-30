package repositories

import (
	"github.com/AntonioFSRE/go-bid/internal/domain/models"
	"github.com/google/uuid"
)

type (
	PGUserRepository interface {
		GetByID(id uuid.UUID) (models.User, error)
		FindByName(name string) (models.User, error)

	}

	RedisUserRepository interface {
		GetByID(id uuid.UUID) (models.User, error)
		GetTokenInfo(id uuid.UUID, tokenID uuid.UUID) (uuid.UUID, error)
		SetToken(id uuid.UUID, tokenID uuid.UUID, exp int64) error
	}
)
