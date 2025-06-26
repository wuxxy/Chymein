package Core

import (
	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo *echo.Echo
	Port string
}

func NewServer(port string) *Server {
	e := echo.New()
	e.Logger.Fatal(e.Start(port))
	return &Server{
		Echo: e,
	}
}
func (s *Server) RegisterRoute(method, path string, handler echo.HandlerFunc) {
	s.Echo.Add(method, path, handler)
}
func (s *Server) RegisterGroup(prefix string, fn func(g *echo.Group)) {
	g := s.Echo.Group(prefix)
	fn(g)
}
