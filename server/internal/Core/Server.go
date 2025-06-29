package Core

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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

var StartTime time.Time

func NewServer(port string) *Server {
	e := echo.New()
	StartTime = time.Now()
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

func LoadConfig() {
	const configPath = "config.json"
	var cfg Cfg

	loadDefaults := func() Cfg {
		return Cfg{
			Port:    Config.Port,
			IsSetup: false,
			Database: DatabaseInfo{
				Host: "localhost",
				Port: "5432",
				User: "",
				Pass: "",
				Name: "",
				SSL:  false,
			},
		}
	}

	writeDefaults := func(cfg Cfg) {
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
		Config = cfg
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
		Config = cfg
		return
	}

	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		fmt.Println("Invalid config JSON, resetting...")
		cfg = loadDefaults()
		writeDefaults(cfg)
		Config = cfg
		return
	}

	Config = cfg
}
func UpdateConfig(updateFn func(cfg *Cfg)) error {
	const configPath = "config.json"

	// Read existing config
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Cfg
	if err := json.Unmarshal(data, &cfg); err != nil {
		return fmt.Errorf("failed to parse config file: %w", err)
	}

	// Apply your update logic
	updateFn(&cfg)

	// Save updated config
	jsonBytes, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal updated config: %w", err)
	}
	if err := os.WriteFile(configPath, jsonBytes, 0644); err != nil {
		return fmt.Errorf("failed to write updated config: %w", err)
	}

	// Update in-memory config too
	Config = cfg

	return nil
}
