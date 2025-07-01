package Database

import (
	"github.com/oklog/ulid/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type Post struct {
	gorm.Model
	ID    string `gorm:"type:char(26);primaryKey"`
	Title string `json:"title"`
	Body  string `json:"body"`
	// Foreign keys
	ContainerID string
	AuthorID    string

	// Associations
	Container  Container         `gorm:"foreignKey:ContainerID"`
	Author     User              `gorm:"foreignKey:AuthorID"`
	Metadata   datatypes.JSONMap `gorm:"type:jsonb"`
	PluginData datatypes.JSONMap `gorm:"type:jsonb"`
}
type Reply struct {
	gorm.Model
	ID   string `gorm:"type:char(26);primaryKey"`
	Body string `json:"body"`
	// Foreign keys
	ParentID string
	AuthorID string

	// Associations
	Parent     Post              `gorm:"foreignKey:ParentID"`
	Author     User              `gorm:"foreignKey:AuthorID"`
	Metadata   datatypes.JSONMap `gorm:"type:jsonb"`
	PluginData datatypes.JSONMap `gorm:"type:jsonb"`
}

func (c *Post) BeforeCreate(_ *gorm.DB) error {
	if c.ID == "" {
		t := time.Now().UTC()
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		c.ID = ulid.MustNew(ulid.Timestamp(t), entropy).String()
	}
	return nil
}

func (c *Reply) BeforeCreate(_ *gorm.DB) error {
	if c.ID == "" {
		t := time.Now().UTC()
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		c.ID = ulid.MustNew(ulid.Timestamp(t), entropy).String()
	}
	return nil
}
