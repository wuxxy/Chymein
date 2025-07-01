package Auth

import (
	"errors"
	"github.com/alexedwards/argon2id"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"server/internal/Core"
	"server/internal/Database"
	"server/internal/User"
	"time"
)

type requestBody struct {
	Username string `json:"username" query:"username" form:"username"`
	Password string `json:"password"  query:"password" form:"password"`
}

func Login(c echo.Context) error {
	var req requestBody
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}
	// Get User
	var user Database.User
	result := Core.DB.Get().First(&user, "username = ?", req.Username)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid username or password",
		})
	} else if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Database error",
		})
	}
	// Check if user locked
	if user.Locked {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Account is locked"})
	}
	// Check password
	isCorrectPassword, err := argon2id.ComparePasswordAndHash(req.Password, user.Password)
	if err != nil {
		log.Println("Password comparison failed:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Auth failure"})
	}

	if !isCorrectPassword {
		if user.LoginAttempts == 3 {
			user.Locked = true
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Account is locked"})
		}
		user.LoginAttempts += 1
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
	}
	if user.Banned {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "You are banned"})
	}
	// Handle login success
	user.LoginAttempts = 0
	user.Locked = false
	user.LastActive = time.Now()
	Core.DB.Get().Save(&user)
	// Create session
	err = User.CreateSession(c, user.ID, 24*time.Hour, "password")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create session",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"login": "success",
	})
}
