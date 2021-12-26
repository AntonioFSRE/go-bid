package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

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
		UserId := int64(0)
		err = rows.Scan(
			&t.BidId,
			&t.Ttl,
			&t.Price,
			&UserId,
			&t.SetAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		t.User = domain.User{
			UserId: UserId,
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *postgresBidRepository) CheckBid(ctx context.Context, bidId int64) (res domain.Bid, err error) {
	query := `SELECT price, ttl, setAt
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

func (m *postgresBidRepository) CreateNewBid(ctx context.Context, bidId int64, ttl int64, price int64) (err error) {
	query := `INSERT  bid SET  bidId=? , ttl=? ,price=? ,setAt=NOW() ,userId=1`
	setAt:=time.Now() 
	userId:=1
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, bidId, ttl, price, setAt, userId)
	if err != nil {
		return
	}
	return
}

func (m *postgresBidRepository) PlaceBid(ctx context.Context, bidId int64, price int64) (err error) {
	query := `UPDATE bid set price=?, userId=? WHERE bidId = ? AND price < ?`
	userId:=2
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return  err
	}

	res, err := stmt.ExecContext(ctx, price, userId)
	if err != nil {
		return err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", affect)
		return err
	}

	return 
}