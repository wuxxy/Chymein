package User

import (
	"context"
	"net/http"
	"server/internal/Core"
	"server/internal/Database"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

const SessionCookieName = "session"
const UserContextKey = "user"

func SessionMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie(SessionCookieName)
			if err != nil || strings.TrimSpace(cookie.Value) == "" {
				return next(c)
			}

			var session Database.Session
			db := Core.DB.Get()
			err = db.First(&session, "id = ? AND revoked = false", cookie.Value).Error
			if err != nil || session.ExpiresAt.Before(time.Now()) {
				if err == nil {
					db.Delete(&session)
				}
				return next(c)
			}

			var user Database.User
			err = db.First(&user, "id = ?", session.UserID).Error
			if err != nil {
				return next(c)
			}

			// Update session's last used time
			session.LastUsedAt = time.Now()
			db.Save(&session)

			ctx := context.WithValue(c.Request().Context(), UserContextKey, &user)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}

func GetUser(c echo.Context) (*Database.User, bool) {
	u := c.Request().Context().Value(UserContextKey)
	user, ok := u.(*Database.User)
	return user, ok
}

func Logout(c echo.Context) error {
	cookie, err := c.Cookie(SessionCookieName)
	if err == nil && strings.TrimSpace(cookie.Value) != "" {
		db := Core.DB.Get()
		db.Delete(&Database.Session{}, "id = ?", cookie.Value)
	}
	c.SetCookie(&http.Cookie{
		Name:     SessionCookieName,
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	return nil
}

func CreateSession(c echo.Context, userID string, duration time.Duration, loginMethod string) error {
	s := &Database.Session{
		UserID:      userID,
		ExpiresAt:   time.Now().Add(duration),
		LastUsedAt:  time.Now(),
		IPAddress:   c.RealIP(),
		UserAgent:   c.Request().UserAgent(),
		LoginMethod: loginMethod,
		Revoked:     false,
	}
	db := Core.DB.Get()
	if err := db.Create(s).Error; err != nil {
		return err
	}
	c.SetCookie(&http.Cookie{
		Name:     SessionCookieName,
		Value:    s.ID,
		Expires:  s.ExpiresAt,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})
	return nil
}
