package User

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
)

func CreateUserHandler(c echo.Context) error {

	return c.String(http.StatusOK, "user endpoint alive")
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}
