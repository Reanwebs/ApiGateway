//go:build wireinject
// +build wireinject

package di

import (
	"gateway/pkg/api"
	"gateway/pkg/api/handlers"
	"gateway/pkg/common/client"
	"gateway/pkg/common/config"

	"github.com/google/wire"
)

// InitializeAPI initializes the API server with dependencies
func InitializeAPI(c *config.Config) (*api.Server, error) {
	wire.Build(
		client.InitClient,
		client.InitConferenceClient,
		client.InitAdminClient,
		client.InitVideoStreamingClient,
		handlers.NewAuthHandler,
		handlers.NewAdminHandler,
		handlers.NewConferenceHandler,
		handlers.NewVideoHandler,
		api.NewServeHTTP,
	)
	return &api.Server{}, nil
}
