package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/Artragnus/go-personal-finance-app/internal/dto"
	"github.com/Artragnus/go-personal-finance-app/internal/entity"
	"github.com/Artragnus/go-personal-finance-app/internal/infra/database"
	"github.com/Artragnus/go-personal-finance-app/internal/token"
)

type HandleUser struct {
	DB        database.UserInterface
	JwtSecret string
}

func NewHandleUser(db database.UserInterface, jwtSecret string) *HandleUser {
	return &HandleUser{
		DB:        db,
		JwtSecret: jwtSecret,
	}
}

func (h *HandleUser) Create(c echo.Context) error {
	var req dto.UserCreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "bad request",
		})
	}

	u, err := h.DB.GetByEmail(req.Email)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": "internal server error",
			})
		}
	}

	if u.ID != uuid.Nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "email already exists",
		})
	}

	user, err := entity.NewUser(req.Name, req.Password, req.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "bad request",
		})
	}

	if err := h.DB.Create(user); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "internal server error",
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *HandleUser) Login(c echo.Context) error {
	var req dto.UserLoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "bad request",
		})
	}

	u, err := h.DB.GetByEmail(req.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "invalid email or password",
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "internal server error",
		})
	}

	if ok := u.ValidatePassword(req.Password); !ok {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid email or password",
		})
	}

	t, err := token.New(u, h.JwtSecret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "internal server error",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
