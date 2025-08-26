package static

import (
	"github.com/gin-gonic/gin"
)

func LoadStaticFiles(router *gin.Engine) {
	router.Static("/assets", "./assets")
}
