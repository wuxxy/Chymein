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
	ID       ulid.ULID `gorm:"type:uuid;primaryKey"`
	Username string    `gorm:"uniqueIndex;not null"`
	Email    string    `gorm:"uniqueIndex;not null"`
	Password string    `gorm:"not null"`

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
	Metadata   datatypes.JSONMap `gorm:"type:jsonb"` // user-specific plugin hooks
	PluginData datatypes.JSONMap `gorm:"type:jsonb"` // plugin-owned state
}

func (u *User) BeforeCreate() (err error) {
	if ulid.ULID.IsZero(u.ID) {
		entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
		ms := ulid.Timestamp(time.Now())
		u.ID, err = ulid.New(ms, entropy)
		if err != nil {
			return err
		}
	}
	return
}

func (u *User) Create(db *gorm.DB) error {
	// Validate required fields
	if u.Username == "" {
		return errors.New("an username is required to create a user")
	}
	if u.Email == "" {
		return errors.New("an email is required to create a user")
	}
	if u.Password == "" {
		return errors.New("a password is required to create a user")
	}

	// Attempt to create the user
	return db.Create(u).Error
}
func (u *User) Update(db *gorm.DB) error {
	// Require a valid ULID (or whatever ID system you're using)
	if ulid.ULID.IsZero(u.ID) {
		return errors.New("user cannot be updated without an ID")
	}

	return db.Model(&User{}).Where("id = ?", u.ID).Updates(u).Error
}
func (u *User) GetByAny(db *gorm.DB) error {
	query := db.Model(&User{})

	switch {
	case u.ID.Compare(ulid.ULID{}) != 0:
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
