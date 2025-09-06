package controllers

import (
	"net/http"
	"strconv"
	"v2/pkg/converters"
	"v2/pkg/errors"
	"v2/pkg/html"
	"v2/pkg/old"
	"v2/pkg/sessions"

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
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/register")
		return
	}
	if controller.userService.CheckUserExists(request.Email) {
		errors.Init()
		errors.Add("Email", "Email already exists")
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))
		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))
	}

	user, err := controller.userService.Create(request)
	if err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}
	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	log.Log.Info("User registered successfully:", user.Name)
	c.Redirect(http.StatusFound, "/")

	c.JSON(http.StatusOK, gin.H{
		"message": "Register done",
	})

}
