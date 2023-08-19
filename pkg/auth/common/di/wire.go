//go:build wireinject
// +build wireinject

package di

import (
	"gateway/pkg/auth/api"
	"gateway/pkg/auth/api/handlers"
	"gateway/pkg/auth/common/client"
	"gateway/pkg/auth/common/config"

	"github.com/google/wire"
)

// InitializeAPI initializes the API server with dependencies
func InitializeAPI(c *config.Config) (*api.Server, error) {
	wire.Build(
		client.InitClient,
		handlers.NewAuthHandler,
		api.NewServeHTTP,
	)
	return &api.Server{}, nil
}
