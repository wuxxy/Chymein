package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"server/internal/Admin"
	"server/internal/Common"
	"server/internal/Core"
	"server/internal/System"
	"server/internal/User"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	LoadConfig()
	StartLogger()
}

func main() {
	server := Core.NewServer(Core.Config.Port)
	server.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	server.RegisterRoute("GET", "/status", Admin.Status)

	if !server.Config.IsSetup {
		server.RegisterRoute("POST", "/setup", Admin.Setup)
	}

	User.RegisterController(server)

	if err := server.Start(); err != nil {
		panic(err)
	}
}

func LoadConfig() {
	// Try opening config

	var cfg Core.Cfg
	var configFile System.File

	_, err := configFile.Open("config.json")

	// Parse content into struct

	if errors.Is(err, System.ErrFileNotFound) {
		fmt.Println("Config file doesn't exist, creating...")

		// Marshal default config struct to JSON
		defaultCfg := Core.Cfg{
			Port:    Core.Config.Port,
			IsSetup: false,
			Database: Core.DatabaseInfo{
				Host: "localhost",
				Port: "5432",
				User: "",
				Pass: "",
				Name: "",
				SSL:  false,
			},
		}
		jsonBytes, err := json.Marshal(defaultCfg)
		if err != nil {
			panic("failed to marshal default config: " + err.Error())
		}

		_, _ = configFile.Create("config.json", jsonBytes, "config")
	}
	_ = configFile.ParseJSON(&cfg)
	Core.Config = cfg
}
func StartLogger() {

	// Format safe filename

	// Create the logs directory if it doesn't exist
	Common.SetLogFile(fmt.Sprintf("logs/%s.txt", Common.StartTime.Format("2006-01-02_15-04-05")))
}
