package client

import (
	"context"
	"gateway/pkg/common/client/interfaces"
	"gateway/pkg/common/config"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AdminClient struct {
	Server pb.AutharizationClient
}

func NewAdminClient(server pb.AutharizationClient) interfaces.AdminClient {
	return &AdminClient{
		Server: server,
	}
}

func InitAdminClient(c *config.Config) (interfaces.AdminClient, error) {
	cc, err := grpc.Dial(c.AuthService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return NewAdminClient(pb.NewAutharizationClient(cc)), nil
}

func (c *AdminClient) AdminLogin(ctx context.Context, request models.AdminLoginRequest) (*pb.AdminLoginResponse, error) {
	res, err := c.Server.AdminLogin(ctx, &pb.AdminLoginRequest{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminClient) GetUsers(ctx context.Context) (*pb.GetUsersResponse, error) {
	res, err := c.Server.GetUsers(ctx, &pb.GetUsersRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminClient) ManageUser(ctx context.Context, request models.ManageUserRequest) (*pb.ManageUserResponse, error) {
	res, err := c.Server.ManageUser(ctx, &pb.ManageUserRequest{
		Id: request.Id,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminClient) GetInterest(ctx context.Context) (*pb.GetInterestResponse, error) {
	res, err := c.Server.GetInterest(ctx, &pb.GetInterestRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminClient) AddInterest(ctx context.Context, request models.AddInterestRequest) (*pb.AddInterestResponse, error) {
	res, err := c.Server.AddInterest(ctx, &pb.AddInterestRequest{
		Interest: request.Interest,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminClient) ManageInterest(ctx context.Context, request models.ManageInterestRequest) (*pb.ManageInterestResponse, error) {
	res, err := c.Server.ManageInterest(ctx, &pb.ManageInterestRequest{
		Id: request.Id,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
