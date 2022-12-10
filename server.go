package main

import (
	"fmt"

	"github.com/khalidrianda/ArticleApp/config"
	"github.com/khalidrianda/ArticleApp/utils/database"
	"github.com/khalidrianda/ArticleApp/utils/helper"

	"github.com/khalidrianda/ArticleApp/features/articles/delivery"
	"github.com/khalidrianda/ArticleApp/features/articles/repository"
	"github.com/khalidrianda/ArticleApp/features/articles/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	db := database.InitDB(cfg)
	database.MigrateDB(db)

	e.Use(helper.Cache().Middleware())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	aRepo := repository.New(db)
	aServices := services.New(aRepo)
	delivery.New(e, aServices)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.ServerPort)))
}
