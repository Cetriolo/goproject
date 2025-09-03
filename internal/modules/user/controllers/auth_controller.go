package controllers

import (
	"net/http"
	"v2/pkg/html"

	"v2/internal/modules/user/requests/auth"
	UserService "v2/internal/modules/user/services"
	log "v2/logger"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService UserService.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		userService: UserService.New(),
	}
}

func (controller *Controller) Register(c *gin.Context) {

	html.Render(c, http.StatusOK, "modules/user/html/register", gin.H{"title": "Register"})

}
func (controller *Controller) HandleRegister(c *gin.Context) {

	var request auth.RegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to bind request: " + err.Error(),
		})
		return
	}

	user, err := controller.userService.Create(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to bind request: " + err.Error(),
		})
		return
	}
	log.Log.Info("User registered successfully:", user.Name)
	c.Redirect(http.StatusFound, "/")

	c.JSON(http.StatusOK, gin.H{
		"message": "Register done",
	})

}
