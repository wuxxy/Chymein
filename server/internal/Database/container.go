package Database

import (
	"github.com/oklog/ulid/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type Container struct {
	gorm.Model
	ID              string `gorm:"type:char(26);primaryKey"`
	Type            string
	Name            string
	Description     string
	ReadPermission  string
	WritePermission string
	ParentID        *string           `gorm:"type:char(26)" json:"parent_id"` // nullable
	Children        []Container       `gorm:"foreignKey:ParentID" json:"children"`
	Metadata        datatypes.JSONMap `gorm:"type:jsonb"`
	PluginData      datatypes.JSONMap `gorm:"type:jsonb"`
	Posts           []Post            `gorm:"foreignKey:ContainerID"`
	SortOrder       int               `gorm:"column:sort_order;default:0"`
}

func (c *Container) BeforeCreate(_ *gorm.DB) error {
	if c.ID == "" {
		t := time.Now().UTC()
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		c.ID = ulid.MustNew(ulid.Timestamp(t), entropy).String()
	}
	return nil
}
