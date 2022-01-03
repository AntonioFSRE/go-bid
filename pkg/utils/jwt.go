package utils

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type JWTConfig struct {
	JWTSecret        string
}

type TokenDetails struct {
	AtID         uuid.UUID
	RtID         uuid.UUID
	AccessToken  string
}

type Claims struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(cfg *JWTConfig, id uuid.UUID) (*TokenDetails, error) {
	atID := uuid.New()
	rtID := uuid.New()


	accessToken, err := createToken(atID, id, cfg.JWTSecret)
	if err != nil {
		return nil, err
	}

	return &TokenDetails{
		AtID:         atID,
		RtID:         rtID,
		AccessToken:  accessToken,
	}, nil
}


func createToken(
	id uuid.UUID,
	userID uuid.UUID,
	secret string,
) (string, error) {
	claims := Claims{
		id.String(),
		userID.String(),
		jwt.StandardClaims{
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(
	c echo.Context,
	tokenName string,
	tokenString string,
	secret string,
) error {
	errString := fmt.Sprintf("invalid %s token", tokenName)

	if tokenString == "" {
		return errors.New(errString)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New(errString)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, ok := claims["id"].(string)
		if !ok {
			return errors.New("invalid jwt claims")
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			return errors.New("invalid jwt claims")
		}

		tokenUuid, err := uuid.Parse(id)
		if err != nil {
			return err
		}

		userUuid, err := uuid.Parse(userID)
		if err != nil {
			return err
		}

		c.Set(fmt.Sprintf("%s_id", tokenName), tokenUuid)
		c.Set("user_id", userUuid)
	}

	return nil
}
