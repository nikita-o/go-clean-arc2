package routes

import (
	"clean-arc-2/internal/handlers/http/controllers"
	"github.com/gin-gonic/gin"
)

func NewApiRouter(gin *gin.Engine, controllers controllers.Controllers) {
	apiRoute := gin.Group("api")
	v1 := apiRoute.Group("v1")

	NewUserRouter(v1, controllers.UserController)

}
