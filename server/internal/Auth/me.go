package Auth

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"server/internal/Database"
	"server/internal/User"
)

func Me(c echo.Context) error {
	user, err := User.GetUser(c)
	if !err {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Couldn't get user",
		})
	}
	if user == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Not logged in",
		})
	}
	return c.JSON(http.StatusOK, Database.User{
		Model:        gorm.Model{},
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		Permissions:  user.Permissions,
		Admin:        user.Admin,
		SuperAdmin:   user.Admin,
		IsVerified:   user.IsVerified,
		Locked:       user.Locked,
		DateOfBirth:  user.DateOfBirth,
		Gender:       user.Gender,
		AvatarURL:    user.AvatarURL,
		BannerColor:  user.BannerColor,
		Bio:          user.Bio,
		PersonalLink: user.PersonalLink,
		Signature:    user.Signature,
		Language:     user.Language,
		Theme:        user.Theme,
		Muted:        user.Muted,
		Metadata:     user.Metadata,
		PluginData:   user.PluginData,
	})
}
