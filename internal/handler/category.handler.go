package handler

import (
	"net/http"

	"github.com/Artragnus/go-personal-finance-app/internal/infra/database"
	"github.com/labstack/echo/v4"
)

type HandleCategory struct {
	DB database.CategoryInterface
}

func NewCategoryHandle(db database.CategoryInterface) *HandleCategory {
	return &HandleCategory{
		DB: db,
	}
}

func (h *HandleCategory) GetIncomeCategories(c echo.Context) error {
	categories, err := h.DB.GetIncomeCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "internal server error",
		})
	}

	return c.JSON(http.StatusOK, categories)
}

func (h *HandleCategory) GetExpenseCategories(c echo.Context) error {
	categories, err := h.DB.GetExpenseCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "internal server error",
		})
	}

	return c.JSON(http.StatusOK, categories)
}
