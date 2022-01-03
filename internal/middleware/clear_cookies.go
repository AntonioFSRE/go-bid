package middleware

import (
	"net/http"

	"github.com/AntonioFSRE/go-bid/internal/config"
	"github.com/AntonioFSRE/go-bid/pkg/logger"
	"github.com/labstack/echo/v4"
)

func ClearCookies(
	cfg *config.Config,
	log logger.Logger,
) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.SetCookie(&http.Cookie{
				Name:   "access_token",
				Path:   "/",
			})

			return next(c)
		}
	}
}
