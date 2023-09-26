package routes

import (
	"gateway/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)

func VideoRoutes(api *gin.RouterGroup, videoHandler handlers.VideoHandler) {
	routes := api.Group("/video")

	// routes.Use(middleware.AuthenticateUser)

	routes.POST("/upload", videoHandler.UploadVideo)
	routes.GET("/user-videos", videoHandler.FetchUserVideo)

}
