package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/Artragnus/go-personal-finance-app/pkg/entity"
)

type payload struct {
	ID entity.ID
	jwt.RegisteredClaims
}

func Auth(secret string) echo.MiddlewareFunc {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(payload)
		},
		SigningKey: []byte(secret),
	}
	return echojwt.WithConfig(config)
}
