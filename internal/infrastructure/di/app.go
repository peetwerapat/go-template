package di

import (
	"github.com/peetwerapat/go-template/internal/infrastructure/repository_impl"
	"github.com/peetwerapat/go-template/internal/usecase"
	"gorm.io/gorm"
)

type AppUseCase struct {
	UserUc *usecase.UserUsecase
}

func InitApp(dbConn *gorm.DB) *AppUseCase {
	userRepo := repository_impl.NewUserRepositoryImplement(dbConn)
	userUC := usecase.NewUserUsecase(userRepo)

	return &AppUseCase{
		UserUc: userUC,
	}
}
