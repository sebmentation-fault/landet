package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sebmentation-fault/landet/views/home"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Static("/static", "assets")

	e.GET("/", func(c echo.Context) error {
		return home.Index().Render(c.Request().Context(), c.Response().Writer)
	}).Name = "Home"

	e.Logger.Fatal(e.Start(":42069"))
}
