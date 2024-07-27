package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sebmentation-fault/landet/views/home"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.File("/favicon.ico", "assets/icons/favicon.ico")
	e.Static("/static", "assets")

	e.GET("/", func(c echo.Context) error {
		return Render(c, http.StatusOK, home.Index())
	}).Name = "Home"

	e.Logger.Fatal(e.Start(":42069"))
}

// NOTE: from the templ docs
// https://github.com/a-h/templ/blob/main/examples/integration-echo/main.go
//
// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
