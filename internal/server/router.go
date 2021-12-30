package server

import (
	authDelivery "github.com/AntonioFSRE/go-bid/internal/auth/delivery/http"
	authRepository "github.com/AntonioFSRE/go-bid/internal/auth/repository"
	authUseCase "github.com/AntonioFSRE/go-bid/internal/auth/usecase"
	bidDelivery "github.com/AntonioFSRE/go-bid/internal/bid/delivery/http"
	bidRepository "github.com/AntonioFSRE/go-bid/internal/bid/repository"
	bidUseCase "github.com/AntonioFSRE/go-bid/internal/bid/usecase"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *Server) middleware() {
	s.router.Pre(middleware.RemoveTrailingSlash())
	s.router.Use(middleware.CORS())
}

func (s *Server) handlers() {
	authRepo := authRepository.NewPGRepository(s.db)
	authRedisRepo := authRepository.NewRedisRepository(s.redis)
	bidRepo := bidRepository.NewPGRepository(s.db)
	bidRedisRepo := bidRepository.NewRedisRepository(s.redis)

	authUC := authUseCase.New(
		s.cfg,
		authRepo,
		authRedisRepo,
		s.log,
	)
	bidUC := bidUseCase.New(
		bidRepo,
		bidRedisRepo,
		s.log,
	)

	if s.cfg.Server.Debug {
		s.router.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	api := s.router.Group("/api")

	authDelivery.Init(
		s.cfg,
		api,
		authUC,
		s.log,
	)
	bidDelivery.Init(
		s.cfg,
		api,
		bidUC,
		authUC,
		s.log,
	)
}
