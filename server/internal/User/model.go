package User

import (
	"errors"
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	// Identity
	ID       string `gorm:"type:char(26);primaryKey"`
	Username string `gorm:"uniqueIndex;not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`

	// Permissions & Roles
	Permissions datatypes.JSONMap `gorm:"type:jsonb"`
	Admin       bool
	SuperAdmin  bool

	// Activity & Auth
	LastActive       time.Time
	LoginAttempts    int  `gorm:"default:0"`
	IsActive         bool `gorm:"default:true"`
	IsVerified       bool `gorm:"default:false"`
	TwoFactorEnabled bool `gorm:"default:false"`
	Sessions         []Session

	// Profile
	DateOfBirth  *time.Time
	Gender       string `gorm:"default:'unselected'"`
	AvatarURL    string
	BannerColor  string `gorm:"default:'brand_color'"`
	Bio          string `gorm:"default:'A newcomer'"`
	PersonalLink string
	Signature    string
	Language     string `gorm:"default:'en_US'"`
	Theme        string `gorm:"default:'light'"`

	// Moderation
	Banned     bool `gorm:"default:false"`
	Muted      bool `gorm:"default:false"`
	AdminNotes string

	// Extensibility
	Metadata   datatypes.JSONMap `gorm:"type:jsonb"`
	PluginData datatypes.JSONMap `gorm:"type:jsonb"`
}

func (u *User) BeforeCreate(_ *gorm.DB) error {
	if u.ID == "" {
		t := time.Now().UTC()
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		u.ID = ulid.MustNew(ulid.Timestamp(t), entropy).String()
	}
	return nil
}

func (u *User) Create(db *gorm.DB) error {
	if u.Username == "" {
		return errors.New("an username is required to create a user")
	}
	if u.Email == "" {
		return errors.New("an email is required to create a user")
	}
	if u.Password == "" {
		return errors.New("a password is required to create a user")
	}
	return db.Create(u).Error
}

func (u *User) Update(db *gorm.DB) error {
	if u.ID == "" {
		return errors.New("user cannot be updated without an ID")
	}
	return db.Model(&User{}).Where("id = ?", u.ID).Updates(u).Error
}

func (u *User) GetByAny(db *gorm.DB) error {
	query := db.Model(&User{})

	switch {
	case u.ID != "":
		query = query.Where("id = ?", u.ID)
	case u.Username != "":
		query = query.Where("username = ?", u.Username)
	case u.Email != "":
		query = query.Where("email = ?", u.Email)
	default:
		return errors.New("no lookup key provided")
	}

	return query.First(u).Error
}

func NewULID() string {
	t := time.Now().UTC()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
