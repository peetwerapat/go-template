package repository

import (
	"github.com/peetwerapat/go-template/internal/domain"
	"github.com/peetwerapat/go-template/internal/interface/repository"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	DB *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) repository.UserRepository {
	return &GormUserRepository{DB: db}
}

func (r *GormUserRepository) CreateUser(user *domain.User) error {
	return r.DB.Create(user).Error
}

func (r *GormUserRepository) GetUserById(id uint) (*domain.User, error) {
	var u domain.User
	if err := r.DB.First(&u, id).Error; err != nil {
		return nil, err
	}

	return &u, nil
}
