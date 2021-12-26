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

func (b *bidUsecase) PlaceBid(c context.Context, u *domain.Bid) (err error) {

	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	return b.bidRepo.PlaceBid(ctx, u)
}

func (b *bidUsecase) CreateNewBid(c context.Context, m *domain.Bid) (err error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()
	existedBid, _ := b.CheckBid(ctx, m.BidId)
	if existedBid != (domain.Bid{}) {
		return domain.ErrConflict
	}

	err = b.bidRepo.CreateNewBid(ctx, m)
	return
}