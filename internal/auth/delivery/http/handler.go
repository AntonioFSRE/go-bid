package http

import (
	"net/http"

	"github.com/AntonioFSRE/go-bid/internal/config"
	"github.com/AntonioFSRE/go-bid/internal/domain/models"
	"github.com/AntonioFSRE/go-bid/internal/domain/usecases"
	"github.com/AntonioFSRE/go-bid/internal/middleware"
	"github.com/AntonioFSRE/go-bid/pkg/logger"
	"github.com/AntonioFSRE/go-bid/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type handler struct {
	cfg         *config.Config
	userUseCase usecases.UserUseCase
	log         logger.Logger
}

func newHandler(
	cfg *config.Config,
	uu usecases.UserUseCase,
	log logger.Logger,
) *handler {
	return &handler{cfg, uu, log}
}

func Init(
	cfg *config.Config,
	e *echo.Group,
	uu usecases.UserUseCase,
	log logger.Logger,
) {
	h := newHandler(cfg, uu, log)
	auth := middleware.Auth(cfg, uu, log)

	authGroup := e.Group("/auth")
	authGroup.POST("/me", h.Me, auth)
	authGroup.POST("/login", h.Login)

	e.GET("/users/:id", h.GetByID)

}


// @Router /auth/me [post]
func (h *handler) Me(c echo.Context) error {
	userID := utils.GetCtxID(c)

	res, err := h.userUseCase.GetByID(userID)
	if err != nil {
		h.log.Errorf("auth.UseCase.GetByID: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, res)
}

// @Router /users/{id} [get]
func (h *handler) GetByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	user, err := h.userUseCase.GetByID(id)
	if err != nil {
		h.log.Errorf("auth.UseCase.GetByID: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

// @Router /auth/login [post]
func (h *handler) Login(c echo.Context) error {
	u := new(models.User)

	if err := c.Bind(u); err != nil {
		return echo.ErrBadRequest
	}

	user, err := h.userUseCase.Login(u)
	if err != nil {
		h.log.Errorf("auth.UseCase.Login: %v", err)
		return err
	}

	h.setCookies(c, user)

	return c.JSON(http.StatusOK, user)
}

func (h *handler) setCookies(c echo.Context, user *models.AuthUser) {
	c.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    user.AccessToken,
		Path:     "/",
		Secure:   h.cfg.Cookie.AccessToken.Secure,
		HttpOnly: h.cfg.Cookie.AccessToken.HttpOnly,
		SameSite: http.SameSiteStrictMode,
	})
}