package User

import (
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
	"math/rand"
	"regexp"
	"time"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}
func NewULID() string {
	t := time.Now().UTC()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
func IsAuthenticated(c echo.Context) bool {
	_, auth := GetUser(c)
	return auth
}
