package Admin

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/internal/Core"
	"server/internal/Database"
	"server/internal/Permission"
	"server/internal/User"
)

const ADMIN_USERS_PERMISSION = "app.admin:users"

func AllUsers(c echo.Context) error {
	user, authenticated := User.GetUser(c)
	if !authenticated {
		return echo.ErrUnauthorized
	}
	if !(Permission.CheckIfUserHasPermission(user, ADMIN_USERS_PERMISSION) || user.SuperAdmin) {
		return echo.ErrUnauthorized
	}
	var users []Database.User

	err := Core.DB.Get().
		Model(&Database.User{}).
		Omit("Password").
		Preload("Roles").
		Preload("Sessions").
		Find(&users).Error
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, users)
}
