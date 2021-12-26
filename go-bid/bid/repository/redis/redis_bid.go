package redis

import (
	"encoding/json"
	"time"

	"github.com/AntonioFSRE/go-bid/bid/repository"
	"github.com/AntonioFSRE/go-bid/domain"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type redisRepository struct {
	redis redis.Store
}

const prefix = "bid"

func NewRedisRepository(rdb redis.Store) repository.RedisArticleRepository {
	return &redisRepository{rdb}
}

func (r *redisRepository) CheckBid(id uuid.UUID) (domain.Bid, error) {
	var bid domain.Bid

	res, err := r.redis.Get(utils.GetRedisKey(prefix, BidId.String()))
	if err != nil {
		return bid,  echo.ErrInternalServerError
	}

	if err := json.Unmarshal([]byte(res), &bid); err != nil {
		return bid, echo.ErrInternalServerError
	}

	return bid, nil
}

func (r *redisRepository) CreateNewBid(
	bid *domain.Bid,
	exp time.Duration,
) error {
	res, err := json.Marshal(bid)
	if err != nil {
		return echo.ErrInternalServerError
	}

	if err := r.redis.Set(utils.GetRedisKey(
		prefix,
		bid.BidId.String(),
	), res, exp); err != nil {
		return echo.ErrInternalServerError
	}

	return nil
}
