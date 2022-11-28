package factory

import (
	authDelivery "14-api-clean-arch/features/auth/delivery"
	authRepo "14-api-clean-arch/features/auth/repository"
	authService "14-api-clean-arch/features/auth/service"

	userDelivery "14-api-clean-arch/features/user/delivery"
	userRepo "14-api-clean-arch/features/user/repository"
	userService "14-api-clean-arch/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userRepoFactory := userRepo.New(db)
	// userRepoFactory := userRepo.NewRaw(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, e)

	authRepoFactory := authRepo.New(db)
	authServiceFactory := authService.New(authRepoFactory)
	authDelivery.New(authServiceFactory, e)

}
