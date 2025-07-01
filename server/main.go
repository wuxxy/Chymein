package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"server/internal/Core"
	"server/internal/User"
	"server/internal/server"
)

func init() {
	Core.LoadConfig()
}

func main() {
	s := Core.NewServer(Core.Config.Port)
	ConnectToDB()
	s.Echo.Use(User.SessionMiddleware())
	s.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true, // 🔥 this is the one you're missing
	}))
	server.Router(s)

	// Router
	// server.RegisterRoute("GET", "/site")
	if err := s.Start(); err != nil {
		panic(err)
	}
}
func ConnectToDB() {

	err := Core.DB.Connect(fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		Core.Config.Database.User,
		Core.Config.Database.Pass,
		Core.Config.Database.Host,
		Core.Config.Database.Port,
		Core.Config.Database.Name,
		func() string {
			if Core.Config.Database.SSL {
				return "require"
			}
			return "disable"
		}(),
	))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Successfully connected to database %s\n", Core.Config.Database.Name)
	}
}
