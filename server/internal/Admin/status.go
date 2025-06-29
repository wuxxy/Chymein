package Admin

import (
	"net/http"
	"server/internal/Core"
	"time"

	"github.com/labstack/echo/v4"
)

func Status(c echo.Context) error {
	status := map[string]interface{}{
		"port":       Core.Config.Port,
		"is_setup":   Core.Config.IsSetup,
		"database":   Core.DB.IsConnected(),
		"time_alive": time.Now().Sub(Core.StartTime).String(),
	}

	return c.JSON(http.StatusOK, status)
}
