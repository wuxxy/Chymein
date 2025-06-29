package User

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	ID     string `gorm:"type:char(26);primaryKey"`
	UserID string `gorm:"type:char(26);not null;index"` // add index for joins
}

// BeforeCreate generates ULID string before insert
func (s *Session) BeforeCreate(_ *gorm.DB) error {
	if s.ID == "" {
		t := time.Now().UTC()
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		s.ID = ulid.MustNew(ulid.Timestamp(t), entropy).String()
	}
	return nil
}
