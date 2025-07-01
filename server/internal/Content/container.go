package Content

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/internal/Core"
	"server/internal/Database"
)

func GetContent(c echo.Context) error {
	var containers []Database.Container

	if err := Core.DB.Get().
		Order(`"sort_order" DESC`).
		Preload("Children").
		Where("parent_id IS NULL").
		Omit("ReadPermissions", "WritePermission", "Metadata", "PluginData", "Posts").
		Find(&containers).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch containers",
		})
	}

	return c.JSON(http.StatusOK, containers)
}
