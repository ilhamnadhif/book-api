package repository

import (
	"book-api/model/schema"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, db *gorm.DB, user schema.User) (schema.User, error)
	FindByEmail(ctx context.Context, db *gorm.DB, email string) (schema.User, error)
}
