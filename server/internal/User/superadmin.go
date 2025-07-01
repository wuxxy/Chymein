package User

import (
	"github.com/alexedwards/argon2id"
	"log"
	"net/http"
	"server/internal/Core"
	"server/internal/Database"
	"time"

	"github.com/labstack/echo/v4"
)

type requestBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateSuperAdmin(c echo.Context) error {
	var req requestBody
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}
	if req.Username == "" || req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "All fields are required",
		})
	}
	if !IsValidEmail(req.Email) {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid email format",
		})
	}
	hash, err := argon2id.CreateHash(req.Password, argon2id.DefaultParams)
	if err != nil {
		log.Println(err)
	}
	adminUser := Database.User{
		ID:               NewULID(),
		Username:         req.Username,
		Email:            req.Email,
		Password:         hash,
		Admin:            true,
		SuperAdmin:       true,
		LastActive:       time.Time{}, // acceptable default
		LoginAttempts:    0,
		IsActive:         false,
		IsVerified:       false,
		TwoFactorEnabled: false,
		Sessions:         nil,          // ← safe if omitted in insert
		DateOfBirth:      nil,          // ← only safe if gorm:"default:null"
		Gender:           "unselected", // ← don't use empty string if you want a defined fallback
		AvatarURL:        "",
		BannerColor:      "brand_color",
		Bio:              "A newcomer",
		PersonalLink:     "",
		Signature:        "",
		Language:         "en_US",
		Theme:            "light",
		Banned:           false,
		Muted:            false,
		AdminNotes:       "",
	}

	session := Database.Session{
		UserID: adminUser.ID,
	}

	if err := Core.DB.Get().Create(&adminUser).Error; err != nil {
		log.Println("Failed to create user:", err)
		return c.String(http.StatusInternalServerError, "Failed to create user")
	}

	// Create session (this will trigger BeforeCreate to generate ULID)
	if err := Core.DB.Get().Create(&session).Error; err != nil {
		log.Println("Failed to create session:", err)
		return c.String(http.StatusInternalServerError, "Failed to create session")
	}

	// Set session cookie using session.ID (already a string)
	c.SetCookie(&http.Cookie{
		Name:     "session",
		Value:    session.ID, // ← already a string
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	// Mark setup complete
	err = Core.UpdateConfig(func(cfg *Core.Cfg) {
		cfg.IsSetup = true
	})
	if err != nil {
		log.Println("Failed to update config:", err)
		return c.String(http.StatusInternalServerError, "Failed to update config")
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"user": adminUser.Username,
	})

}
