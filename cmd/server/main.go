package main

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/Artragnus/go-personal-finance-app/configs"
	"github.com/Artragnus/go-personal-finance-app/internal/entity"
	"github.com/Artragnus/go-personal-finance-app/internal/handler"
	"github.com/Artragnus/go-personal-finance-app/internal/infra/database"
	"github.com/Artragnus/go-personal-finance-app/internal/middleware"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)

	db := database.New(dsn)
	db.AutoMigrate(
		&entity.User{},
		entity.Expense{},
		entity.Income{},
		entity.CategoryExpense{},
		entity.CategoryIncome{},
	)

	database.Seed(db)

	userDb := database.NewUser(db)
	userHandler := handler.NewHandleUser(userDb, config.JWTSecret)

	expenseDb := database.NewExpense(db)
	expenseHandler := handler.NewHandleExpense(expenseDb)

	incomeDb := database.NewIncome(db)
	incomeHandler := handler.NewHandleIncome(incomeDb)

	categoryDb := database.NewCategory(db)
	categoryHandler := handler.NewCategoryHandle(categoryDb)

	e := echo.New()
	e.POST("/user", userHandler.Create)
	e.POST("/login", userHandler.Login)

	userMiddleware := middleware.User(config.JWTSecret)

	expense := e.Group("expense", userMiddleware)
	expense.POST("", expenseHandler.Create)
	expense.GET("", expenseHandler.Get)

	income := e.Group("income", userMiddleware)
	income.POST("", incomeHandler.Create)
	income.GET("", incomeHandler.Get)

	category := e.Group("category", userMiddleware)
	category.GET("/income", categoryHandler.GetIncomeCategories)
	category.GET("/expense", categoryHandler.GetExpenseCategories)

	err = e.Start(config.WebServerPort)
	if err != nil {
		panic(err)
	}
}
