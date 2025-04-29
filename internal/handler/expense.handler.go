package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Artragnus/go-personal-finance-app/internal/dto"
	"github.com/Artragnus/go-personal-finance-app/internal/entity"
	"github.com/Artragnus/go-personal-finance-app/internal/infra/database"
	"github.com/Artragnus/go-personal-finance-app/internal/token"
)

type HandleExpense struct {
	DB database.ExpenseInterface
}

func NewHandleExpense(db database.ExpenseInterface) *HandleExpense {
	return &HandleExpense{
		DB: db,
	}
}

func (h *HandleExpense) Create(c echo.Context) error {
	var req dto.ExpenseCreateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "bad request",
		})
	}

	payload := token.GetPayload(c)

	expense, err := entity.NewExpense(payload.ID, req.Amount, req.Description, req.CategoryID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	res, err := h.DB.Create(expense)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "internal server error",
		})
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *HandleExpense) Get(c echo.Context) error {
	payload := token.GetPayload(c)

	res, err := h.DB.Get(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, res)
}
