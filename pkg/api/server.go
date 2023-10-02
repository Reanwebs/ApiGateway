package api

import (
	"gateway/pkg/api/handlers"
	"gateway/pkg/api/routes"
	"gateway/pkg/common/config"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	Engine *gin.Engine
	Port   string
}

func NewServeHTTP(c *config.Config, userHandler handlers.UserHandler, conferenceHandler handlers.ConferenceHandler,
	adminHandler handlers.AdminHandler, videoHandler handlers.VideoHandler, monitaization handlers.MonitizationHandler) (*Server, error) {

	engine := gin.New()
	engine.Use(gin.Logger())

	engine.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.UserRoutes(engine.Group("/api"), userHandler, videoHandler)
	routes.AdminRoutes(engine.Group("/api"), adminHandler, videoHandler)
	routes.ConferenceRoutes(engine.Group("/api"), conferenceHandler)
	routes.VideoRoutes(engine.Group("/api"), videoHandler)
	routes.MonitaizationRoutes(engine.Group("/api"), monitaization)
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"StatusCode": 404,
			"msg":        "invalid url",
		})
	})

	return &Server{
		Engine: engine,
		Port:   c.Port,
	}, nil
}

func (c *Server) Start() {
	c.Engine.Run(c.Port)
}
