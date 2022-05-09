package repository

import (
	"book-api/model/schema"
	"context"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, db *gorm.DB, user schema.User) (schema.User, error) {

	res := db.WithContext(ctx).Create(&user)
	err := res.Error
	if err != nil {
		db.Rollback()
		return user, err
	}
	return user, nil
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, db *gorm.DB, email string) (schema.User, error) {
	var user schema.User
	res := db.WithContext(ctx).First(&user, "email = ?", email)
	err := res.Error
	if err != nil {
		db.Rollback()
		return user, err
	}
	return user, nil
}
