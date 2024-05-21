package main

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	group := e.Group("/docs")
	group.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper: func(c echo.Context) bool {
			return strings.TrimSuffix(c.Request().URL.Path, "/") == "/docs"
		},
		Root:       "docs",
		Browse:     true,
		IgnoreBase: false,
	}))
	e.Start(":3015")
}
