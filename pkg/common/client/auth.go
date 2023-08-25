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
		CPassword:   request.CPassword,
		PhoneNumber: request.PhoneNumber,
		Username:    request.UserName,
		Referral:    request.Referral,
		Otp:         request.Otp,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) OtpRequest(ctx context.Context, request models.OtpValidation) (*pb.OtpSignUpResponse, error) {
	res, err := c.Server.OtpRequest(ctx, &pb.OtpSignUpRequest{
		Username:    request.UserName,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Password:    request.Password,
		CPassword:   request.CPassword,
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

func (c *authClient) ResendOtp(ctx context.Context, request models.ResendOtp) (*pb.ResendOtpResponse, error) {
	res, err := c.Server.ResendOtp(ctx, &pb.ResendOtpRequest{
		PhoneNumber: request.PhoneNumber,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
