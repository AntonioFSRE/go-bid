package models

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	User struct {
		ID        uuid.UUID `json:"id" db:"id" example:"1"`
		Name      string    `json:"name" db:"name" validate:"required,name" example:"admin"`
		Password  string    `json:"password,omitempty" db:"password" validate:"omitempty,min=6,max=250" swaggerignore:"true"`
		Role      string    `json:"role,omitempty" db:"role" validate:"required,role" example:"admin"`
	}

	AuthUser struct {
		User         *User  `json:"user"`
		TokenType    string `json:"token_type" validate:"required" example:"Bearer"`
		AccessToken  string `json:"access_token" validate:"required"`
	}
)

func (u *User) Validate() error {
	validate := validator.New()

	u.Name = strings.ToLower(strings.TrimSpace(u.Name))
	u.Password = strings.TrimSpace(u.Password)

	return validate.Struct(u)
}

func (u *User) ValidatePassword() error {
	if u.Password == "" {
		return errors.New("empty password")
	}

	return nil
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(u.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil
}

func (u *User) ComparePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword(
		[]byte(u.Password),
		[]byte(password),
	); err != nil {
		return err
	}

	return nil
}

func (u *User) SanitizePassword() {
	u.Password = ""
}
