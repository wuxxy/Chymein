package User

import (
	"net/http"
	"server/internal/Core"

	"github.com/labstack/echo/v4"
)

func RegisterController(server *Core.Server) {
	server.RegisterRoute("POST", "/api/user/create", createUserHandler)
}

func createUserHandler(c echo.Context) error {

	return c.String(http.StatusOK, "user endpoint alive")
}
