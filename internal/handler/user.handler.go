package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Artragnus/go-personal-finance-app/internal/dto"
	"github.com/Artragnus/go-personal-finance-app/internal/entity"
	"github.com/Artragnus/go-personal-finance-app/internal/infra/database"
)

type HandleUser struct {
	DB database.UserInterface
}

func NewHandleUser(db database.UserInterface) *HandleUser {
	return &HandleUser{
		DB: db,
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
		log.Println(err)
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
