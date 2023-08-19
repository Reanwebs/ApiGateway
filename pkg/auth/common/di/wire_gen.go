// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"gateway/pkg/auth/api"
	"gateway/pkg/auth/api/handlers"
	"gateway/pkg/auth/common/client"
	"gateway/pkg/auth/common/config"
)

// Injectors from wire.go:

// InitializeAPI initializes the API server with dependencies

func InitializeAPI(c *config.Config) (*api.Server, error) {
	authClient, err := client.InitClient(c)
	if err != nil {
		return nil, err
	}
	userHandler := handlers.NewAuthHandler(authClient)
	server, err := api.NewServeHTTP(c, userHandler)
	if err != nil {
		return nil, err
	}
	return server, nil
}
