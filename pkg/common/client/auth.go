package client

import (
	"context"
	"errors"
	"fmt"
	"gateway/pkg/common/client/interfaces"
	"gateway/pkg/common/config"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb"
	"gateway/pkg/utils"
	"time"

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

// func (c *authClient) UserLogin(ctx context.Context, request models.LoginRequestBody) (*pb.LoginResponse, error) {
// 	res, err := c.Server.UserLogin(ctx, &pb.LoginRequest{
// 		Email:    request.Email,
// 		Password: request.Password,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil

// }

func (c *authClient) UserLogin(ctx context.Context, request models.LoginRequestBody, retryConfig models.RetryConfig) (*pb.LoginResponse, error) {
	var res *pb.LoginResponse
	var err error

	startTime := time.Now()

	for retryCount := 1; retryCount <= retryConfig.MaxRetries; retryCount++ {
		fmt.Println("try.........", retryCount)

		if time.Since(startTime) > retryConfig.MaxDuration {
			err = errors.New("time limit exceeded")
			break
		}

		ctx1, cancel := context.WithTimeout(ctx, 30*time.Second)
		done := make(chan struct{})

		go func() {
			defer close(done)
			<-ctx.Done()
		}()

		select {
		case <-done:
			cancel()
			return nil, context.Canceled
		default:
			res, err = c.Server.UserLogin(ctx1, &pb.LoginRequest{
				Email:    request.Email,
				Password: request.Password,
			})
			cancel()

			if err == nil {
				return res, nil
			}

			// Check if the error is non-retryable
			if utils.IsNonRetryableError(err) {
				return nil, err
			}
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryConfig.RetryInterval):
		}
	}

	return nil, err
}

// func (c *authClient) UserLogin(ctx context.Context, request models.LoginRequestBody) (*pb.LoginResponse, error) {
// 	var res *pb.LoginResponse
// 	var err error
// 	maxRetries := 2
// 	for i := 0; i <= maxRetries; i++ {
// 		fmt.Println("try\t", i)
// 		ctx1, cancel := context.WithTimeout(ctx, 30*time.Second)
// 		defer cancel()

// 		done := make(chan struct{})
// 		go func() {
// 			defer close(done)
// 			<-ctx.Done() // Listen for cancellation
// 		}()

// 		select {
// 		case <-done:
// 			// Handle cancellation
// 			return nil, context.Canceled
// 		default:
// 			res, err = c.Server.UserLogin(ctx1, &pb.LoginRequest{
// 				Email:    request.Email,
// 				Password: request.Password,
// 			})
// 			if err == nil {
// 				return res, nil // Request succeeded, return response
// 			}
// 			// Request failed, continue to retry if possible
// 		}

// 		select {
// 		case <-ctx.Done():
// 			// Handle timeout
// 			return nil, context.DeadlineExceeded
// 		case <-time.After(time.Second): // Wait before retrying
// 		}
// 	}

// 	return nil, err // Return last error after maxRetries
// }

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

func (c *authClient) ForgotPasswordOtp(ctx context.Context, request models.ForgotPasswordOtpRequest) (*pb.ForgotPasswordOtpResponse, error) {
	res, err := c.Server.ForgotPasswordOtp(ctx, &pb.ForgotPasswordOtpRequest{
		PhoneNumber: request.PhoneNumber,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) ForgotPasswordValidateOtp(ctx context.Context, request models.ForgotPasswordValidateOtpRequest) (*pb.ForgotPasswordValidateOtpResponse, error) {
	res, err := c.Server.ForgotPasswordValidateOtp(ctx, &pb.ForgotPasswordValidateOtpRequest{
		PhoneNumber: request.PhoneNumber,
		Otp:         request.Otp,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) ForgotPasswordChangePassword(ctx context.Context, request models.ForgotPasswordChangePasswordRequest) (*pb.ForgotPasswordChangePasswordResponse, error) {
	res, err := c.Server.ForgotPasswordChangePassword(ctx, &pb.ForgotPasswordChangePasswordRequest{
		PhoneNumber: request.PhoneNumber,
		Password:    request.Password,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) ValidateUser(ctx context.Context, request models.ValidateUserRequest) (*pb.ValidateUserResponse, error) {
	res, err := c.Server.ValidateUser(ctx, &pb.ValidateUserRequest{
		Email: request.Email,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) GoogleLogin(ctx context.Context, request models.GoogleLoginRequest) (*pb.GoogleLoginResponse, error) {
	res, err := c.Server.GoogleLogin(ctx, &pb.GoogleLoginRequest{
		Token: request.Token,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
