package usecase

import (
	"time"

	"github.com/AntonioFSRE/go-bid/internal/domain/models"
	"github.com/AntonioFSRE/go-bid/internal/domain/repositories"
	"github.com/AntonioFSRE/go-bid/internal/domain/usecases"
	"github.com/AntonioFSRE/go-bid/pkg/logger"
)

type usecase struct {
	pgRepository    repositories.PGBidRepository
	redisRepository repositories.RedisBidRepository
	log             logger.Logger
}

const cacheDuration = 3600

func New(
	pg repositories.PGBidRepository,
	redis repositories.RedisBidRepository,
	log logger.Logger,
) usecases.BidUseCase {
	return &usecase{
		pgRepository:    pg,
		redisRepository: redis,
		log:             log,
	}
}


func (u *usecase) GetByID(id int64) (models.Bid, error) {
	cachedBid, err := u.redisRepository.GetByID(id)
	if err != nil {
		u.log.Errorf("bid.redisRepository.GetByID: %v", err)
	}

	if cachedBid.ID != 0 {
		return cachedBid, nil
	}

	res, err := u.pgRepository.GetByID(id)
	if err != nil {
		u.log.Errorf("bid.pgRepository.GetByID: %v", err)
		return res, err
	}

	if err := u.redisRepository.SetBid(
		&res,
		time.Second*cacheDuration,
	); err != nil {
		u.log.Errorf("bid.redisRepository.SetBid: %v", err)
		return res, err
	}

	return res, nil
}

func (u *usecase) Store(bid *models.Bid) (*models.Bid, error) {

	res, err := u.pgRepository.Store(bid)
	if err != nil {
		u.log.Errorf("bid.pgRepository.Store: %v", err)
		return nil, err
	}

	if err := u.redisRepository.SetBid(
		res,
		time.Second*cacheDuration,
	); err != nil {
		u.log.Errorf("bid.redisRepository.SetBid: %v", err)
		return nil, err
	}

	return res, nil
}

func (u *usecase) Update(bid *models.Bid) (*models.Bid, error) {

	res, err := u.pgRepository.Update(bid)
	if err != nil {
		u.log.Errorf("bid.pgRepository.Update: %v", err)
		return nil, err
	}

	if err := u.redisRepository.SetBid(
		res,
		time.Second*cacheDuration,
	); err != nil {
		u.log.Errorf("bid.redisRepository.SetBid: %v", err)
		return nil, err
	}

	return res, nil
}