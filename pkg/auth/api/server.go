package api

import (
	"gateway/pkg/auth/api/handlers"
	"gateway/pkg/auth/api/routes"
	"gateway/pkg/auth/common/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
	Port   string
}

func NewServeHTTP(c *config.Config, userHandler handlers.UserHandler) (*Server, error) {
	engine := gin.New()
	engine.Use(gin.Logger())

	routes.UserRoutes(engine.Group("/user"), userHandler)
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
