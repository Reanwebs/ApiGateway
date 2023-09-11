package api

import (
	"fmt"
	"gateway/pkg/api/handlers"
	"gateway/pkg/api/routes"
	"gateway/pkg/common/config"
	"net/http"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	Engine *gin.Engine
	Port   string
	Socket *socketio.Server
}

func NewServeHTTP(c *config.Config, userHandler handlers.UserHandler,
	conferenceHandler handlers.ConferenceHandler, adminHandler handlers.AdminHandler, videoHandler handlers.VideoHandler) (*Server, error) {

	engine := gin.New()
	engine.Use(gin.Logger())

	server := socketio.NewServer(nil)

	engine.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.UserRoutes(engine.Group("/api"), userHandler)
	routes.AdminRoutes(engine.Group("/api"), adminHandler)
	routes.ConferenceRoutes(engine.Group("/api"), conferenceHandler)
	routes.VideoRoutes(engine.Group("/api"), videoHandler)
	fmt.Println("socket-IO connected")
	// Connect WebSocket server
	engine.GET("/socket.io", gin.WrapH(server))

	server.OnConnect("/webRTCPeers", func(s socketio.Conn) error {

		clientID := s.ID()

		fmt.Println("Client connected:", clientID)
		return nil
	})

	server.OnEvent("/webRTCPeers", "chat-message", func(s socketio.Conn, msg string) {
		s.Emit("chat-message", msg)
		fmt.Println(msg)
	})

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
