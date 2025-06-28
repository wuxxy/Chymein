package Core

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Cfg struct {
	Port     string       `json:"port"`
	Database DatabaseInfo `json:"database"`
	IsSetup  bool         `json:"is_setup"`
}

type Server struct {
	Echo   *echo.Echo
	Config Cfg
}

func NewServer(port string) *Server {
	e := echo.New()
	return &Server{
		Echo: e,
		Config: Cfg{
			Port: port,
		},
	}
}

func (s *Server) RegisterRoute(method, path string, handler echo.HandlerFunc) {
	fmt.Printf("[+] Registering (%s) %s\n", method, path)
	s.Echo.Add(method, path, handler)
}
func (s *Server) RegisterGroup(prefix string, fn func(g *echo.Group)) {
	fmt.Printf("[+] Registering Group (%s)\n", prefix)
	group := s.Echo.Group(prefix)
	fn(group)
}

func (s *Server) Start() error {
	return s.Echo.Start(":" + s.Config.Port)
}

// WriteCookie sets a cookie with the given name, value, and expiration duration.
func (s *Server) WriteCookie(c echo.Context, name, value string, duration time.Duration) {
	cookie := &http.Cookie{
		Name:    name,
		Value:   value,
		Expires: time.Now().Add(duration),
	}
	c.SetCookie(cookie)
}

// ReadCookie returns the cookie with the given name, or an error if not found.
func (s *Server) ReadCookie(c echo.Context, name string) (*http.Cookie, error) {
	return c.Cookie(name)
}

// ReadAllCookies returns all cookies in the current request.
func (s *Server) ReadAllCookies(c echo.Context) []*http.Cookie {
	return c.Cookies()
}

var Config Cfg = Cfg{}
