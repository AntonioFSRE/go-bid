package repository

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/AntonioFSRE/go-bid/internal/domain/models"
	"github.com/AntonioFSRE/go-bid/internal/domain/repositories"
	"github.com/AntonioFSRE/go-bid/pkg/store/redis"
	"github.com/AntonioFSRE/go-bid/pkg/utils"
	"github.com/labstack/echo/v4"
)

type redisRepository struct {
	redis redis.Store
}

const prefix = "bid"

func NewRedisRepository(rdb redis.Store) repositories.RedisBidRepository {
	return &redisRepository{rdb}
}

func (r *redisRepository) GetByID(id int64) (models.Bid, error) {
	var bid models.Bid

	res, err := r.redis.Get(utils.GetRedisKey(prefix, strconv.FormatInt(id,10)))
	if err != nil {
		return bid, echo.ErrNotFound
	}

	if err := json.Unmarshal([]byte(res), &bid); err != nil {
		return bid, echo.ErrInternalServerError
	}

	return bid, nil
}

func (r *redisRepository) SetBid(
	bid *models.Bid,
	exp time.Duration,
) error {
	res, err := json.Marshal(bid)
	if err != nil {
		return echo.ErrInternalServerError
	}

	if err := r.redis.Set(utils.GetRedisKey(
		prefix,
		strconv.FormatInt(bid.ID,10),
	), res, exp); err != nil {
		return echo.ErrInternalServerError
	}

	return nil
}
