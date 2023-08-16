package routes

import (
	"gateway/pkg/auth/api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterVideoRoutes(api *gin.RouterGroup, userHandler handlers.UserHandler) {

	routes := api.Group("/user")
	routes.POST("/register", userHandler.UserSignup)
	routes.POST("/login", userHandler.UserLogin)

}
