package Admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Setup(c echo.Context) error {
	return c.String(http.StatusOK, ":)")
}
