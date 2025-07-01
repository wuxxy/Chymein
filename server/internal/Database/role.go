package Database

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID          uint `gorm:"primaryKey;autoIncrement:false"`
	Name        string
	Permissions []string `gorm:"type:text[]"`
	Color       string
}
