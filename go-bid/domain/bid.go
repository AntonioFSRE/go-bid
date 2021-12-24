package domain

import (
	"context"
	"time"
)

// Bid ...
type Bid struct {
	bidId     int64     `json:"bidid"`
	ttl       int64     `json:"ttl"`
	price     int64     `json:"price"`
	setAt     time.Time `json:"set_at"`
	User      User      `json:"user"`
}

// BidUsecase represent the bid's usecases
type BidUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]Bid, string, error)
	CreateNewBid(context.Context, *Bid) error
	CheckBid(ctx context.Context, bidId int64) (Bid, error)
	PlaceBid(ctx context.Context, u *Bid) error
}

// BidRepository represent the bid's repository contract
type BidRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Bid, nextCursor string, err error)
	CreateNewBid(ctx context.Context, b *Bid) error
	CheckBid(ctx context.Context, bidId int64) (Bid, error)
	PlaceBid(ctx context.Context, u *Bid) error
}