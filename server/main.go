package main

import (
	"fmt"
	"server/internal/Admin"
	"server/internal/Core"
	"server/internal/User"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	Core.LoadConfig()
}

func main() {
	server := Core.NewServer(Core.Config.Port)
	ConnectToDB()
	server.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	server.RegisterRoute("GET", "/status", Admin.Status)

	if !server.Config.IsSetup {
		server.RegisterRoute("POST", "/create_admin", Admin.CreateSuperAdmin)
	}

	server.RegisterRoute("POST", "/api/user/create", User.CreateUserHandler)

	// Router
	server.RegisterRoute("GET", "/site")

	if err := server.Start(); err != nil {
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
