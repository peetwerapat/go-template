package repository

import (
	"github.com/peetwerapat/go-template/internal/domain"
)

type UserRepository interface {
	GetUserById(id uint) (*domain.User, error)
	CreateUser(user *domain.User) error
}
