package main

import (
	"fmt"

	"github.com/jabeedhexanovamedia/todo-ap/config"
	"github.com/labstack/echo/v5"

	"github.com/labstack/echo/v5/middleware"
)

func main() {
	cfg := config.LoadConfig()

	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.GET("/", func(c *echo.Context) error {
		return c.String(200, "Todo API running in "+cfg.AppEnv+" mode")
	})

	port := fmt.Sprintf(":%s", cfg.Port)
	if err := e.Start(port); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}

}
