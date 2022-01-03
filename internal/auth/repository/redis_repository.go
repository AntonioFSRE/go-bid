package repository

import (
	"encoding/json"

	"github.com/AntonioFSRE/go-bid/internal/domain/models"
	"github.com/AntonioFSRE/go-bid/internal/domain/repositories"
	"github.com/AntonioFSRE/go-bid/pkg/store/redis"
	"github.com/AntonioFSRE/go-bid/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type redisRepository struct {
	redis redis.Store
}

const (
	authPrefix = "auth"
	userPrefix = "users"
)

func NewRedisRepository(rdb redis.Store) repositories.RedisUserRepository {
	return &redisRepository{rdb}
}

func (r *redisRepository) GetByID(id uuid.UUID) (models.User, error) {
	var user models.User

	res, err := r.redis.Get(utils.GetRedisKey(userPrefix, id.String()))
	if err != nil {
		return user, echo.ErrNotFound
	}

	if err = json.Unmarshal([]byte(res), &user); err != nil {
		return user, echo.ErrInternalServerError
	}

	return user, nil
}

func (r *redisRepository) GetTokenInfo(
	id uuid.UUID,
	tokenID uuid.UUID,
) (uuid.UUID, error) {
	res, err := r.redis.Get(utils.GetRedisKey(
		authPrefix,
		id.String(),
		tokenID.String(),
	))
	if err != nil {
		return uuid.Nil, echo.ErrNotFound
	}

	return uuid.Parse(res)
}

func (r *redisRepository) SetToken(
	id uuid.UUID,
	tokenID uuid.UUID,
) error {

	if err := r.redis.Set(utils.GetRedisKey(
		authPrefix,
		id.String(),
		tokenID.String(),
	), id.String()); err != nil {
		return echo.ErrInternalServerError
	}

	return nil
}