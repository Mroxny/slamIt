package repository

import (
	"context"

	"github.com/Mroxny/slamIt/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	*Repository[model.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Repository: NewRepository[model.User](db),
	}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return &user, err
}
