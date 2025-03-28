package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"github.com/Artragnus/go-personal-finance-app/internal/entity"
	pkgEntity "github.com/Artragnus/go-personal-finance-app/pkg/entity"
)

type payload struct {
	ID    pkgEntity.ID
	Email string
	jwt.RegisteredClaims
}

func New(u entity.User, secret string) (string, error) {
	payload := &payload{
		u.ID,
		u.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return t, err
	}

	return t, nil
}

func GetPayload(c echo.Context) *payload {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*payload)

	payload := &payload{
		Email: claims.Email,
		ID:    claims.ID,
	}

	return payload
}
