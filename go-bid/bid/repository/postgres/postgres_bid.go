package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/AntonioFSRE/go-bid/bid/repository"
	"github.com/AntonioFSRE/go-bid/domain"
	"github.com/sirupsen/logrus"
)

type postgresBidRepository struct {
	Conn *sql.DB
}

// NewPostgresBidRepository will create an object that represent the bid.Repository interface
func NewPostgresBidRepository(Conn *sql.DB) domain.BidRepository {
	return &postgresBidRepository{Conn}
}

func (m *postgresBidRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Bid, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domain.Bid, 0)
	for rows.Next() {
		t := domain.Bid{}
		userId := int64(0)
		err = rows.Scan(
			&t.bidId,
			&t.ttl,
			&t.price,
			&userId,
			&t.setAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		t.User = domain.User{
			userId: userId,
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *postgresBidRepository) Fetch(ctx context.Context, cursor string, num int64) (res []domain.Bid, nextCursor string, err error) {
	query := `SELECT bidId,ttl,price, userId, setAt
  						FROM bid`

	decodedCursor, err := repository.DecodeCursor(cursor)
	if err != nil && cursor != "" {
		return nil, "", domain.ErrBadParamInput
	}

	res, err = m.fetch(ctx, query, decodedCursor, num)
	if err != nil {
		return nil, "", err
	}

	if len(res) == int(num) {
		nextCursor = repository.EncodeCursor(res[len(res)-1].setAt)
	}

	return
}

func (m *postgresBidRepository) CheckBid(ctx context.Context, bidId int64) (res domain.Bid, err error) {
	query := `SELECT bidId,ttl,price
  						FROM bid WHERE bidId = ?`

	list, err := m.fetch(ctx, query, bidId)
	if err != nil {
		return domain.Bid{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return 
}

func (m *postgresBidRepository) CreateNewBid(ctx context.Context, b *domain.Bid) (err error) {
	query := `INSERT  bid SET  bidId=? , ttl=? ,price=? ,setAt=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, b.bidId, b.ttl, b.price, time.Now())
	if err != nil {
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	b.bidId = lastID
	return
}

func (m *postgresBidRepository) PlaceBid(ctx context.Context, u *domain.Bid) (err error) {
	query := `UPDATE bid set price=? WHERE bidId = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, u.price, u.bidId)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", affect)
		return
	}

	return
}