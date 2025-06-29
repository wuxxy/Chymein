package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"server/internal/Admin"
	"server/internal/Core"
	"server/internal/User"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	LoadConfig()
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
		server.RegisterRoute("POST", "/setup", Admin.Setup)
	}

	User.RegisterController(server)

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
func LoadConfig() {
	const configPath = "config.json"
	var cfg Core.Cfg

	loadDefaults := func() Core.Cfg {
		return Core.Cfg{
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
	}

	writeDefaults := func(cfg Core.Cfg) {
		jsonBytes, err := json.MarshalIndent(cfg, "", "  ")
		if err != nil {
			panic("failed to marshal default config: " + err.Error())
		}
		err = os.WriteFile(configPath, jsonBytes, 0644)
		if err != nil {
			panic("failed to write default config: " + err.Error())
		}
		fmt.Println("Default config written to", configPath)
	}

	// Try to open the config file
	file, err := os.OpenFile(configPath, os.O_RDWR, 0644)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Config file not found, creating with defaults.")
		cfg = loadDefaults()
		writeDefaults(cfg)
		Core.Config = cfg
		return
	} else if err != nil {
		panic("failed to open config file: " + err.Error())
	}
	defer file.Close()

	// Try to parse the config
	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Failed to read config, resetting...")
		cfg = loadDefaults()
		writeDefaults(cfg)
		Core.Config = cfg
		return
	}

	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		fmt.Println("Invalid config JSON, resetting...")
		cfg = loadDefaults()
		writeDefaults(cfg)
		Core.Config = cfg
		return
	}

	Core.Config = cfg
}
