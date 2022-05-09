package schema

import (
	"time"
)

type User struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"not null;type:varchar(100)"`
	Email     string `gorm:"unique;not null;type:varchar(100)"`
	Password  string `gorm:"not null;type:varchar(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Books     []Book `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
