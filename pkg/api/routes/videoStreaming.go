package routes

import (
	"gateway/pkg/api/handlers"
	middleware "gateway/pkg/common/midleware"

	"github.com/gin-gonic/gin"
)

func VideoRoutes(api *gin.RouterGroup, videoHandler handlers.VideoHandler) {
	routes := api.Group("/video")

	routes.Use(middleware.AuthenticateUser)

	routes.POST("/upload", videoHandler.UploadVideo)

}
