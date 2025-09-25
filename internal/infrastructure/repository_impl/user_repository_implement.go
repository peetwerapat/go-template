package repository_impl

import (
	"github.com/peetwerapat/go-template/internal/domain"
	"github.com/peetwerapat/go-template/internal/interface/repository"
	"gorm.io/gorm"
)

type UserRepositoryImplement struct {
	DB *gorm.DB
}

func NewUserRepositoryImplement(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryImplement{DB: db}
}

func (r *UserRepositoryImplement) CreateUser(user *domain.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepositoryImplement) GetUserById(id uint) (*domain.User, error) {
	var u domain.User
	if err := r.DB.First(&u, id).Error; err != nil {
		return nil, err
	}

	return &u, nil
}
