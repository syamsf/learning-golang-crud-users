package routes

import (
	userController "syamsf/learning-gin/src/modules/user/controller"
	userService "syamsf/learning-gin/src/modules/user/service"
	userRepo "syamsf/learning-gin/src/modules/user/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func Api(router *gin.Engine, db *gorm.DB) {
	validate := validator.New()
	userRepository := userRepo.NewUserRepository(db)
	userService 	 := userService.NewUserService(userRepository, validate)
	userController := userController.NewUserControllerImpl(userService)

	v1 := router.Group("/api/v1") 
	
	users := v1.Group("/users")
	{
		users.GET("", userController.FindAll)
		users.POST("", userController.Create)
		users.GET(":id", userController.FindById)
		users.PATCH(":id", userController.Update)
		users.DELETE(":id", userController.Delete)
	}
}