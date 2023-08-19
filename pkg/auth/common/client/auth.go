package client

import (
	"context"
	"gateway/pkg/auth/common/client/interfaces"
	"gateway/pkg/auth/common/config"
	"gateway/pkg/auth/common/models"
	"gateway/pkg/auth/common/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type authClient struct {
	Server pb.AutharizationClient
}

func InitClient(c *config.Config) (interfaces.AuthClient, error) {
	cc, err := grpc.Dial(c.AuthService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return NewAuthClient(pb.NewAutharizationClient(cc)), nil
}

func NewAuthClient(server pb.AutharizationClient) interfaces.AuthClient {
	return &authClient{
		Server: server,
	}
}

func (c *authClient) UserSignup(ctx context.Context, request models.RegisterRequestBody) (*pb.SignupResponse, error) {

	res, err := c.Server.UserSignup(ctx, &pb.SignupRequest{
		Email:       request.Email,
		Password:    request.Password,
		PhoneNumber: request.PhoneNumber,
		UserName:    request.UserName,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) OtpValidation(ctx context.Context, request models.OtpValidation) (*pb.OtpValidationResponse, error) {
	res, err := c.Server.OtpValidation(ctx, &pb.OtpValidationRequest{
		Otp:         request.Otp,
		Phonenumber: request.PhoneNumber,
	})
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (c *authClient) UserLogin(ctx context.Context, request models.LoginRequestBody) (*pb.LoginResponse, error) {
	res, err := c.Server.UserLogin(ctx, &pb.LoginRequest{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (c *authClient) ValidName(ctx context.Context, request models.ValidName) (*pb.ValidNameResponse, error) {
	res, err := c.Server.ValidName(ctx, &pb.ValidNameRequest{
		Username: request.UserName,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) ValidName(ctx context.Context, request models.ValidName) (*pb.ValidNameResponse, error) {
	res, err := c.Server.ValidName(ctx, &pb.ValidNameRequest{
		Username: request.UserName,
	})
	if err != nil {
		return nil, err
	}
	return res, nil

}

func NewAuthClient(server pb.AutharizationClient) interfaces.AuthClient {
	return &authClient{
		Server: server,
	}
}

func (c *authClient) UserSignup(ctx context.Context, request models.RegisterRequestBody) (*pb.SignupResponse, error) {

	res, err := c.Server.UserSignup(ctx, &pb.SignupRequest{
		Email:       request.Email,
		Password:    request.Password,
		PhoneNumber: request.PhoneNumber,
		UserName:    request.UserName,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) OtpValidation(ctx context.Context, request models.OtpValidation) (*pb.OtpValidationResponse, error) {
	res, err := c.Server.OtpValidation(ctx, &pb.OtpValidationRequest{
		Otp:         request.Otp,
		Phonenumber: request.PhoneNumber,
	})
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (c *authClient) UserLogin(ctx context.Context, request models.LoginRequestBody) (*pb.LoginResponse, error) {
	res, err := c.Server.UserLogin(ctx, &pb.LoginRequest{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (c *authClient) ValidName(ctx context.Context, request models.ValidName) (*pb.ValidNameResponse, error) {
	res, err := c.Server.ValidName(ctx, &pb.ValidNameRequest{
		Username: request.UserName,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
