package domain

import (
	"context"
	"time"
)

// Bid ...
type Bid struct {
	BidId     int64     `json:"bidid"`
	Ttl       int64     `json:"ttl"`
	Price     int64     `json:"price"`
	SetAt     time.Time `json:"set_at"`
	User      User      `json:"user"`
}

// BidUsecase represent the bid's usecases
type BidUsecase interface {
	CreateNewBid(ctx context.Context, bidId int64, ttl int64, price int64) error
	CheckBid(ctx context.Context, bidId int64) (Bid, error)
	PlaceBid(ctx context.Context, bidId int64, price int64) error 
}

// BidRepository represent the bid's repository contract
type BidRepository interface {
	CreateNewBid(ctx context.Context, bidId int64, ttl int64, price int64) error
	CheckBid(ctx context.Context, bidId int64) (Bid, error)
	PlaceBid(ctx context.Context, bidId int64, price int64) (error)
}