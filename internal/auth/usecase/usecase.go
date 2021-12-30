package usecase

import (
	"net/http"

	"github.com/AntonioFSRE/go-bid/internal/config"
	"github.com/AntonioFSRE/go-bid/internal/domain/models"
	"github.com/AntonioFSRE/go-bid/internal/domain/repositories"
	"github.com/AntonioFSRE/go-bid/internal/domain/usecases"
	"github.com/AntonioFSRE/go-bid/pkg/logger"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type usecase struct {
	cfg             *config.Config
	pgRepository    repositories.PGUserRepository
	redisRepository repositories.RedisUserRepository
	log             logger.Logger
}

const cacheDuration = 3600

func New(
	cfg *config.Config,
	pg repositories.PGUserRepository,
	redis repositories.RedisUserRepository,
	log logger.Logger,
) usecases.UserUseCase {
	return &usecase{
		cfg:             cfg,
		pgRepository:    pg,
		redisRepository: redis,
		log:             log,
	}
}


func (u *usecase) GetByID(id uuid.UUID) (models.User, error) {
	cachedUser, err := u.redisRepository.GetByID(id)
	if err != nil {
		u.log.Errorf("auth.redisRepository.GetByID: %v", err)
	}

	if cachedUser.ID != uuid.Nil {
		return cachedUser, nil
	}

	res, err := u.pgRepository.GetByID(id)
	if err != nil {
		u.log.Errorf("auth.pgRepository.GetByID: %v", err)
		return res, err
	}

	res.SanitizePassword()

	return res, nil
}

func (u *usecase) Login(user *models.User) (*models.AuthUser, error) {
	if err := user.Validate(); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := user.ValidatePassword(); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := u.pgRepository.FindByName(user.Name)
	if err != nil {
		u.log.Errorf("auth.pgRepository.FindByName: %v", err)
		return nil, err
	}

	if err = res.ComparePassword(user.Password); err != nil {
		return nil, echo.ErrUnauthorized
	}

	res.SanitizePassword()

	return u.Auth(&res)
}