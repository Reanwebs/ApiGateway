package api

import (
	"gateway/pkg/api/handlers"
	"gateway/pkg/api/routes"
	"gateway/pkg/common/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
	Port   string
}

func NewServeHTTP(c *config.Config, userHandler handlers.UserHandler, conferenceHandler handlers.ConferenceHandler) (*Server, error) {
	engine := gin.New()
	engine.Use(gin.Logger())

	routes.UserRoutes(engine.Group("/api"), userHandler, conferenceHandler)
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
