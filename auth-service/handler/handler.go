package handler

import (
	"errors"

	gt "github.com/golang-jwt/jwt"
)

var (
	ErrInvalidToken = errors.New("invalid token provided")
	ErrForbidden    = errors.New("resource forbidden")
	DefaultSecret   = "howh3ogf4q34hf9ogq34f"
)

type authClaims struct {
	Roles []string `json:"roles"`
	Name  string   `json:"name"`
	gt.StandardClaims
}

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}
