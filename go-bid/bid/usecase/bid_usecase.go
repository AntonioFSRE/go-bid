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

// NewBidUsecase will create new BidUsecase object representation of domain.BidUsecase interface
func NewBidUsecase(b domain.BidRepository, u domain.UserRepository, timeout time.Duration) domain.BidUsecase {
	return &bidUsecase{
		bidRepo:    b,
		userRepo:   u,
		contextTimeout: timeout,
	}
}

func (b *bidUsecase) Fetch(c context.Context, cursor string, num int64) (res []domain.Bid, nextCursor string, err error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	res, nextCursor, err = b.bidRepo.Fetch(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}

	return
}

func (b *bidUsecase) CheckBid(c context.Context, id int64) (res domain.Bid, err error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	res, err = b.bidRepo.CheckBid(ctx, bidId)
	if err != nil {
		return
	}

	resUser, err := u.userRepo.CheckBid(ctx, res.User.userId)
	if err != nil {
		return domain.User{}, err
	}
	res.User = resUser
	return
}

func (b *bidUsecase) PlaceBid(c context.Context, u *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()

	return b.bidRepo.PlaceBid(ctx, u)
}

func (b *bidUsecase) CreateNewBid(c context.Context, m *domain.Bid) (err error) {
	ctx, cancel := context.WithTimeout(c, b.contextTimeout)
	defer cancel()
	existedBid, _ := b.CheckBid(ctx, m.bidId)
	if existedBid != (domain.Bid{}) {
		return domain.ErrConflict
	}

	err = b.BidRepo.Store(ctx, m)
	return
}

