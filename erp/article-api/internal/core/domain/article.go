package domain

import (
	"time"

	"gorm.io/gorm"
)

// Article descripe the basic article Model ..
type Article struct {
	gorm.Model
	ID          uint
	ArticleID   uint32 `gorm:"primaryKey"`
	Title       string
	Description string
	Group       string
	Typ         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type ArticleListItem struct {
	ID          uint
	ArticleID   uint32 `gorm:"primaryKey"`
	Title       string
	Description string
	Group       string
	Typ         string
	SupplierID  string
	OwnerID     string
}
