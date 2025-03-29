package middleware

import (
	"github.com/golang-jwt/jwt/v5"

	"github.com/Artragnus/go-personal-finance-app/pkg/entity"
)

type payload struct {
	ID entity.ID
	jwt.RegisteredClaims
}
