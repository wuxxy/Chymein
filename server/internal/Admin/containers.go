package Admin

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/internal/Core"
	"server/internal/Database"
	"server/internal/Permission"
	"server/internal/User"
)

const ADMIN_CONTAINERS_PERMISSION = "app.admin:containers"

type createContainerReq struct {
	Name        string `json:"name" query:"name" form:"name"`
	Description string `json:"description"  query:"description" form:"description"`
	Type        string `json:"type"  query:"type" form:"type"`
	ReadPerm    string `json:"read_perm"  query:"read_perm" form:"read_perm"`
	WritePerm   string `json:"write_perm"  query:"write_perm" form:"write_perm"`
	ParentID    string `json:"parent_id"  query:"parent_id" form:"parent_id"`
}

func AllContainers(c echo.Context) error {
	user, authenticated := User.GetUser(c)
	if !authenticated {
		return echo.ErrUnauthorized
	}
	if !(Permission.CheckIfUserHasPermission(user, ADMIN_CONTAINERS_PERMISSION) || user.SuperAdmin) {
		return echo.ErrUnauthorized
	}
	var containers []Database.Container

	if err := Core.DB.Get().
		Order(`"sort_order" DESC`).
		Preload("Children").
		Where("parent_id IS NULL").
		Find(&containers).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch containers",
		})
	}

	return c.JSON(http.StatusOK, containers)
}

func CreateContainer(c echo.Context) error {
	var req createContainerReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	if req.Name == "" || req.Type == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "All fields are required",
		})
	}

	user, authenticated := User.GetUser(c)
	if !authenticated {
		return echo.ErrUnauthorized
	}

	if !(Permission.CheckIfUserHasPermission(user, ADMIN_CONTAINERS_PERMISSION) || user.SuperAdmin) {
		return echo.ErrUnauthorized
	}

	container := Database.Container{
		Type:            req.Type,
		Name:            req.Name,
		Description:     req.Description,
		ReadPermission:  req.ReadPerm,
		WritePermission: req.WritePerm,
	}

	// Check if a parent ID is provided and assign it
	if req.ParentID != "" {
		var parent Database.Container
		tx := Core.DB.Get().First(&parent, "id = ?", req.ParentID)
		if tx.Error != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Parent container not found",
			})
		}
		container.ParentID = &parent.ID
	}

	if err := Core.DB.Get().Create(&container).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create container",
		})
	}

	return c.JSON(http.StatusCreated, container)
}
