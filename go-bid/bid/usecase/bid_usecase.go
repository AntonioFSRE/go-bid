package usecase

import (
	"context"
	"time"

	"github.com/AntonioFSRE/go-bid/domain"
)

type bidUsecase struct {
	bidRepo    domain.BidRepository
	userRepo     domain.UserRepository
	contextTimeout time.Duration
}

func NewBidUsecase(b domain.BidRepository, u domain.UserRepository, timeout time.Duration) domain.BidUsecase {
	return &bidUsecase{
		bidRepo:    b,
		userRepo:   u,
		contextTimeout: timeout,
	}
}

func (b *bidUsecase) CheckBid(c context.Context, bidId int64) (res domain.Bid, err error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	res, err = b.bidRepo.CheckBid(ctx, bidId)
	if err != nil {
		return
	}

	resUser, err := b.userRepo.GetByID(ctx, res.User.UserId)
	if err != nil {
		return domain.Bid{}, err
	}

	res.User = resUser
	return
}

func (b *bidUsecase) PlaceBid(c context.Context, bidId int64, price int64) (err error) {

	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	return b.bidRepo.PlaceBid(ctx, bidId, price)
}

func (b *bidUsecase) CreateNewBid(c context.Context, bidId int64, ttl int64, price int64) (err error) {


	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()
	existedBid, _ := b.CheckBid(ctx, bidId)
	if existedBid != (domain.Bid{}) {
		return domain.ErrConflict
	}

	err = b.bidRepo.CreateNewBid(ctx, bidId, ttl, price)
	return
}