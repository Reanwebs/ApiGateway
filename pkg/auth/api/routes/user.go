package routes


import (
	"gateway/pkg/auth/api/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup, userHandler handlers.UserHandler) {
	routes := api.Group("/user")
	routes.POST("/signup", userHandler.UserSignup)
	routes.POST("/otp", userHandler.OtpValidation)
	routes.POST("/login", userHandler.UserLogin)
}

