package routes

import (
	homeController "v2/internal/modules/home/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	homeController := homeController.New()

	router.GET("/", homeController.Index)

}
