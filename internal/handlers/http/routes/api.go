package routes

import (
	"clean-arc-2/internal/handlers/http/controllers"
	"clean-arc-2/internal/handlers/http/middlewares"
	"github.com/gin-gonic/gin"
)

func NewApiRouter(gin *gin.Engine, controllers *controllers.Controllers, middlewares *middlewares.Middlewares) {
	apiRoute := gin.Group("api")
	versionApiRoute := apiRoute.Group(":version")

	NewUserRouter(versionApiRoute, controllers.UserController)
}
