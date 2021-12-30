package usecase

import (
	"github.com/AntonioFSRE/go-bid/internal/domain/models"
	"github.com/AntonioFSRE/go-bid/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const authPrefix = "auth"

func (u *usecase) Auth(user *models.User) (*models.AuthUser, error) {
	td, err := utils.GenerateToken(
		&utils.JWTConfig{
			JWTSecret:        u.cfg.Server.JwtSecret,
			JWTRefreshSecret: u.cfg.Server.JwtRefreshSecret,
			AtExpires:        u.cfg.Cookie.AccessToken.MaxAge,
			RtExpires:        u.cfg.Cookie.RefreshToken.MaxAge,
		},
		user.ID,
	)
	if err != nil {
		u.log.Errorf("generateToken: %v", err)
		return nil, echo.ErrInternalServerError
	}

	if err := u.redisRepository.SetToken(
		user.ID,
		td.AtID,
		td.AtExpires,
	); err != nil {
		u.log.Errorf("auth.redisRepository.SetToken: %v", err)
		return nil, err
	}

	if err := u.redisRepository.SetToken(
		user.ID,
		td.RtID,
		td.RtExpires,
	); err != nil {
		u.log.Errorf("auth.redisRepository.SetToken: %v", err)
		return nil, err
	}

	return &models.AuthUser{
		User:         user,
		TokenType:    "Bearer",
		ExpiresIn:    u.cfg.Cookie.AccessToken.MaxAge,
		AccessToken:  td.AccessToken,
		RefreshToken: td.RefreshToken,
	}, nil
}

func (u *usecase) GetToken(
	id uuid.UUID,
	tokenID uuid.UUID,
) (uuid.UUID, error) {
	id, err := u.redisRepository.GetTokenInfo(id, tokenID)
	if err != nil {
		u.log.Errorf("auth.redisRepository.GetTokenInfo: %v", err)
		return uuid.Nil, err
	}

	return id, nil
}

