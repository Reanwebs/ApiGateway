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
	routes.GET("/user-videos", videoHandler.FetchUserVideo)
	routes.GET("/user-archived-videos", videoHandler.FetchUserArchivedVideo)
	routes.GET("/list-all", videoHandler.FetchAllVideo)
	routes.POST("/archive-video", videoHandler.ArchivVideo)
	routes.GET("/get-by-id", videoHandler.GetVideoById)
	routes.PATCH("/star", videoHandler.ToggleStar)
	routes.PUT("/report-video", videoHandler.ReportVideo)
	routes.GET("/exclusive-content", videoHandler.FetchExclusiveVideo)

	adminroutes := api.Group("/video")
	adminroutes.Use(middleware.AuthenticateAdmin)

	adminroutes.PATCH("/block-video", videoHandler.BlockVideo)
	adminroutes.GET("/get-reported-videos", videoHandler.GetReportedVideos)

}
