package repository

import (
	"github.com/peetwerapat/go-template/internal/domain"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserById(id uint) (*domain.User, error)
}
