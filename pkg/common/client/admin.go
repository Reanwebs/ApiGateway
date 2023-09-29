package client

import (
	"context"
	"gateway/pkg/common/client/interfaces"
	"gateway/pkg/common/config"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AdminClient struct {
	Server auth.AutharizationClient
}

func NewAdminClient(server auth.AutharizationClient) interfaces.AdminClient {
	return &AdminClient{
		Server: server,
	}
}

func InitAdminClient(c *config.Config) (interfaces.AdminClient, error) {
	cc, err := grpc.Dial(c.AuthService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return NewAdminClient(auth.NewAutharizationClient(cc)), nil
}

func (c *AdminClient) AdminLogin(ctx context.Context, request models.AdminLoginRequest) (*auth.AdminLoginResponse, error) {
	res, err := c.Server.AdminLogin(ctx, &auth.AdminLoginRequest{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminClient) GetUsers(ctx context.Context) (*auth.GetUsersResponse, error) {
	res, err := c.Server.GetUsers(ctx, &auth.GetUsersRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminClient) ManageUser(ctx context.Context, request models.ManageUserRequest) (*auth.ManageUserResponse, error) {
	res, err := c.Server.ManageUser(ctx, &auth.ManageUserRequest{
		Id: request.Id,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminClient) GetInterest(ctx context.Context) (*auth.GetInterestResponse, error) {
	res, err := c.Server.GetInterest(ctx, &auth.GetInterestRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminClient) AddInterest(ctx context.Context, request models.AddInterestRequest) (*auth.AddInterestResponse, error) {
	res, err := c.Server.AddInterest(ctx, &auth.AddInterestRequest{
		Interest: request.Interest,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminClient) ManageInterest(ctx context.Context, request models.ManageInterestRequest) (*auth.ManageInterestResponse, error) {
	res, err := c.Server.ManageInterest(ctx, &auth.ManageInterestRequest{
		Id: request.Id,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminClient) ManageCommunity(ctx context.Context, request models.ManageCommunityRequest) (*auth.ManageCommunityResponse, error) {
	res, err := c.Server.ManageCommunity(ctx, &auth.ManageCommunityRequest{
		CommunityId: request.CommunityId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminClient) GetAllCommunity(ctx context.Context) (*auth.GetAllCommunityResponse, error) {
	res, err := c.Server.GetAllCommunity(ctx, &auth.GetAllCommunityRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}
