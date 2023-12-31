// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"gateway/pkg/api"
	"gateway/pkg/api/handlers"
	"gateway/pkg/common/client"
	"gateway/pkg/common/config"
)

// Injectors from wire.go:

// InitializeAPI initializes the API server with dependencies
func InitializeAPI(c *config.Config) (*api.Server, error) {
	authClient, err := client.InitClient(c)
	if err != nil {
		return nil, err
	}
	userHandler := handlers.NewAuthHandler(authClient)
	conferenceClient, err := client.InitConferenceClient(c)
	if err != nil {
		return nil, err
	}
	conferenceHandler := handlers.NewConferenceHandler(conferenceClient)
	adminClient, err := client.InitAdminClient(c)
	if err != nil {
		return nil, err
	}
	adminHandler := handlers.NewAdminHandler(adminClient)
	videoClient, err := client.InitVideoStreamingClient(c)
	if err != nil {
		return nil, err
	}
	videoHandler := handlers.NewVideoHandler(videoClient)
	monitaizationClient, err := client.InitMonitizationClient(c)
	if err != nil {
		return nil, err
	}
	monitizationHandler := handlers.NewMonitizationHandler(monitaizationClient)
	server, err := api.NewServeHTTP(c, userHandler, conferenceHandler, adminHandler, videoHandler, monitizationHandler)
	if err != nil {
		return nil, err
	}
	return server, nil
}
