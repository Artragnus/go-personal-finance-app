package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Artragnus/go-personal-finance-app/internal/dto"
	"github.com/Artragnus/go-personal-finance-app/internal/entity"
	"github.com/Artragnus/go-personal-finance-app/internal/infra/database"
	"github.com/Artragnus/go-personal-finance-app/internal/token"
)

type HandleIncome struct {
	DB database.IncomeInterface
}

func NewHandleIncome(db database.IncomeInterface) *HandleIncome {
	return &HandleIncome{
		DB: db,
	}
}

func (h *HandleIncome) Create(c echo.Context) error {
	var req dto.IncomeCreateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "bad request",
		})
	}

	payload := token.GetPayload(c)

	income, err := entity.NewIncome(payload.ID, req.Amount, req.Description, req.CategoryID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	res, err := h.DB.Create(income)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "internal server error",
		})
	}

	return c.JSON(http.StatusCreated, res)
}
