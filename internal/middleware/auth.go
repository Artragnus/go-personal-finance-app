package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/Artragnus/go-personal-finance-app/internal/token"
)

func User(secret string) echo.MiddlewareFunc {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(token.Payload)
		},
		SigningKey: []byte(secret),
	}
	return echojwt.WithConfig(config)
}
