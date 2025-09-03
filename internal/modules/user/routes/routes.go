package routes

import (
	userController "v2/internal/modules/user/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	userController := userController.New()

	router.GET("/register", userController.Register)
	router.POST("/register", userController.HandleRegister)

}
