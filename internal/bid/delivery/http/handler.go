package http

import (
	"net/http"
	"strconv"

	"github.com/AntonioFSRE/go-bid/internal/config"
	"github.com/AntonioFSRE/go-bid/internal/domain/models"
	"github.com/AntonioFSRE/go-bid/internal/domain/usecases"
	"github.com/AntonioFSRE/go-bid/internal/middleware"
	"github.com/AntonioFSRE/go-bid/pkg/logger"
	"github.com/AntonioFSRE/go-bid/pkg/utils"
	"github.com/labstack/echo/v4"
)

type handler struct {
	bidUseCase usecases.BidUseCase
	userUseCase    usecases.UserUseCase
	log            logger.Logger
}

func newHandler(
	au usecases.BidUseCase,
	uu usecases.UserUseCase,
	log logger.Logger,
) *handler {
	return &handler{
		bidUseCase: au,
		userUseCase:    uu,
		log:            log,
	}
}

func Init(
	cfg *config.Config,
	e *echo.Group,
	au usecases.BidUseCase,
	uu usecases.UserUseCase,
	log logger.Logger,
) {
	h := newHandler(au, uu, log)
	auth := middleware.Auth(cfg, uu, log)

	e.GET("/check/:id", h.GetByID)
	e.POST("/lock/:id/:ttl/:price", h.Store, auth)
	e.PUT("/bid/:id/:price", h.Update, auth)
}

// @Router /check/{id} [get]
func (h *handler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	b_id := int64(id)
	if err != nil {
		return echo.ErrNotFound
	}

	bid, err := h.bidUseCase.GetByID(b_id)
	if err != nil {
		h.log.Errorf("bid.UseCase.GetByID: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, bid)
}


// @Router /lock/{id}/{ttl}/{price} [post]
func (h *handler) Store(c echo.Context) error {
	a := new(models.Bid)

	if err := c.Bind(a); err != nil {
		return echo.ErrBadRequest
	}

	a.AuthorID = utils.GetCtxID(c)

	createdBid, err := h.bidUseCase.Store(a)
	if err != nil {
		h.log.Errorf("bid.UseCase.Store: %v", err)
		return err
	}

	return c.JSON(http.StatusCreated, createdBid)
}


// @Router /bid/{id}/{price} [put]
func (h *handler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	a := new(models.Bid)

	if err := c.Bind(a); err != nil {
		return echo.ErrBadRequest
	}

	a.ID = int64(id)
	a.AuthorID = utils.GetCtxID(c)

	updatedBid, err := h.bidUseCase.Update(a)
	if err != nil {
		h.log.Errorf("bid.UseCase.Update: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, updatedBid)
}