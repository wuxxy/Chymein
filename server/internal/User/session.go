package User

import (
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	ID     ulid.ULID `gorm:"primary_key"`
	UserID uint
}
