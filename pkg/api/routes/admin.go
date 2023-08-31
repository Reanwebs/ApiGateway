package routes

import (
	"gateway/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(api *gin.RouterGroup, adminHandler handlers.AdminHandler) {
	routes := api.Group("/admin")
	routes.POST("login", adminHandler.AdminLogin)
	routes.GET("get-users", adminHandler.GetUsers)
	routes.PATCH("manage-users", adminHandler.ManageUser)
	routes.GET("get-intrest", adminHandler.GetInterest)
	routes.POST("add-intrest", adminHandler.AddInterest)
}
