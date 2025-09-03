package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/peetwerapat/go-template/docs"
	"github.com/peetwerapat/go-template/internal/infrastructure/di"
	"github.com/peetwerapat/go-template/internal/interface/controller"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(appUc *di.AppUseCase) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// CORS config
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PACTH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// Auth Controller
	authController := controller.NewAuthController(appUc.UserUc)
	r.POST("/register", authController.CreateUser)

	return r
}
