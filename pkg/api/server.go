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

func NewServeHTTP(c *config.Config, userHandler handlers.UserHandler, conferenceHandler handlers.ConferenceHandler) (*Server, error) {
	engine := gin.New()
	engine.Use(gin.Logger())

	engine.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.UserRoutes(engine.Group("/api"), userHandler)
	routes.ConferenceRoutes(engine.Group("/api"), conferenceHandler)

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
