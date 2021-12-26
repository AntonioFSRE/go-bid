package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/AntonioFSRE/go-bid/domain"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// BidHandler  represent the httphandler for article
type BidHandler struct {
	BUsecase domain.BidUsecase
}

// NewBidHandler will initialize the bids/ resources endpoint
func NewBidHandler(e *echo.Echo, us domain.BidUsecase) {
	handler := &BidHandler{
		BUsecase: us,
	}
	e.POST("/lock/:bidId/:ttl/:price", handler.CreateNewBid)
	e.GET("/check/:bidId", handler.CheckBid)
	e.PUT("/bid/:bidId/:price", handler.PlaceBid)
}

// CheckBid will get bid by given id
func (b *BidHandler) CheckBid(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("bidId"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	bidId := int64(idP)
	ctx := c.Request().Context()

	bd, err := b.BUsecase.CheckBid(ctx, bidId)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, bd)
}

func isRequestValid(m *domain.Bid) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateNewBid will store the bid by given request body
func (b *BidHandler) CreateNewBid(c echo.Context) (err error) {
	var bid domain.Bid
	err = c.Bind(&bid)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&bid); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = b.BUsecase.CreateNewBid(ctx, &bid)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, bid)
}

func (b *BidHandler) PlaceBid(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("bidId"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	a := new(domain.Bid)

	if err := c.Bind(a); err != nil {
		return echo.ErrNotFound
	}

	a.BidId = int64(idP)

	ctx := c.Request().Context()
	err = b.BUsecase.PlaceBid(ctx, a)
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	return c.JSON(http.StatusOK, a)
}


func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}