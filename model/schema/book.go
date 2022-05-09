package schema

import (
	"time"
)

type Book struct {
	ID          int    `gorm:"primaryKey"`
	UserID      int    `gorm:"not null;type:integer"`
	Name        string `gorm:"unique;not null;type:varchar(100)"`
	Description string `gorm:"type:text"`
	Price       int    `gorm:"not null;type:integer"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
