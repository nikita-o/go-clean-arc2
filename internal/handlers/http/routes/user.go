package routes

import (
	"clean-arc-2/internal/handlers/http/controllers"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(router *gin.RouterGroup, uc *controllers.UserController) {
	userRoute := router.Group("user")

	userRoute.GET("", uc.GetUserById)
	userRoute.PATCH("", uc.UpdateUser)
	userRoute.DELETE("", uc.DeleteUser)
	userRoute.POST("", uc.CreateUser)
}
