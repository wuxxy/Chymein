package main

import (
	"net/http"
	"server/internal/Core"

	"github.com/labstack/echo/v4"
)

func main() {
	e := Core.NewServer("4569").Echo
	e.GET("/h_w", func(c echo.Context) error {
		return c.String(http.StatusOK, ":)")
	})

}
