package main

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/Artragnus/go-personal-finance-app/configs"
	"github.com/Artragnus/go-personal-finance-app/internal/entity"
	"github.com/Artragnus/go-personal-finance-app/internal/handler"
	"github.com/Artragnus/go-personal-finance-app/internal/infra/database"
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
	db.AutoMigrate(&entity.User{})

	userDb := database.NewUser(db)
	userHandler := handler.NewHandleUser(userDb)

	e := echo.New()

	e.POST("/user", userHandler.Create)

	err = e.Start(config.WebServerPort)
	if err != nil {
		panic(err)
	}
}
