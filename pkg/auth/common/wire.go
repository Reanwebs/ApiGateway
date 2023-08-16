package common

import (
	"gateway/pkg/auth/api"
	"gateway/pkg/auth/api/handlers"
	"gateway/pkg/auth/common"
	"gateway/pkg/auth/common/client"

	"github.com/google/wire"
)

func InitializeAPI(c *common.Config) (*api.Server, error) {
	wire.Build(client.InitServiceClient, handlers.NewAuthHandler, api.NewServeHTTP)
	return &api.Server{}, nil
}
