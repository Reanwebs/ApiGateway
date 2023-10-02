package client

import (
	"context"
	"errors"
	"fmt"
	"gateway/pkg/common/client/interfaces"
	"gateway/pkg/common/config"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb/auth"
	"gateway/pkg/utils"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type authClient struct {
	Server auth.AutharizationClient
}

func InitClient(c *config.Config) (interfaces.AuthClient, error) {
	cc, err := grpc.Dial(c.AuthService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return NewAuthClient(auth.NewAutharizationClient(cc)), nil
}

func NewAuthClient(server auth.AutharizationClient) interfaces.AuthClient {
	return &authClient{
		Server: server,
	}
}

func (c *authClient) UserSignup(ctx context.Context, request models.RegisterRequestBody, retryConfig models.RetryConfig) (*auth.SignupResponse, error) {

	res, err := c.Server.UserSignup(ctx, &auth.SignupRequest{
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

func (c *authClient) OtpRequest(ctx context.Context, request models.OtpValidation, retryConfig models.RetryConfig) (*auth.OtpSignUpResponse, error) {
	res, err := c.Server.OtpRequest(ctx, &auth.OtpSignUpRequest{
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

func (c *authClient) UserLogin(ctx context.Context, request models.LoginRequestBody, retryConfig models.RetryConfig) (*auth.LoginResponse, error) {
	var res *auth.LoginResponse
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
			res, err = c.Server.UserLogin(ctx1, &auth.LoginRequest{
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

func (c *authClient) ValidName(ctx context.Context, request models.ValidName, retryConfig models.RetryConfig) (*auth.ValidNameResponse, error) {
	var res *auth.ValidNameResponse
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
			res, err = c.Server.ValidName(ctx1, &auth.ValidNameRequest{
				Username: request.UserName,
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

func (c *authClient) ResendOtp(ctx context.Context, request models.ResendOtp, retryConfig models.RetryConfig) (*auth.ResendOtpResponse, error) {

	var res *auth.ResendOtpResponse
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
			res, err = c.Server.ResendOtp(ctx1, &auth.ResendOtpRequest{
				PhoneNumber: request.PhoneNumber,
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

func (c *authClient) ForgotPasswordOtp(ctx context.Context, request models.ForgotPasswordOtpRequest, retryConfig models.RetryConfig) (*auth.ForgotPasswordOtpResponse, error) {

	var res *auth.ForgotPasswordOtpResponse
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
			res, err = c.Server.ForgotPasswordOtp(ctx1, &auth.ForgotPasswordOtpRequest{
				PhoneNumber: request.PhoneNumber,
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

func (c *authClient) ForgotPasswordValidateOtp(ctx context.Context, request models.ForgotPasswordValidateOtpRequest, retryConfig models.RetryConfig) (*auth.ForgotPasswordValidateOtpResponse, error) {

	var res *auth.ForgotPasswordValidateOtpResponse
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
			res, err = c.Server.ForgotPasswordValidateOtp(ctx1, &auth.ForgotPasswordValidateOtpRequest{
				PhoneNumber: request.PhoneNumber,
				Otp:         request.Otp,
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

func (c *authClient) ForgotPasswordChangePassword(ctx context.Context, request models.ForgotPasswordChangePasswordRequest, retryConfig models.RetryConfig) (*auth.ForgotPasswordChangePasswordResponse, error) {

	var res *auth.ForgotPasswordChangePasswordResponse
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
			res, err = c.Server.ForgotPasswordChangePassword(ctx1, &auth.ForgotPasswordChangePasswordRequest{
				PhoneNumber: request.PhoneNumber,
				Password:    request.Password,
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

func (c *authClient) ValidateUser(ctx context.Context, request models.ValidateUserRequest, retryConfig models.RetryConfig) (*auth.ValidateUserResponse, error) {

	var res *auth.ValidateUserResponse
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
			res, err = c.Server.ValidateUser(ctx1, &auth.ValidateUserRequest{
				Email: request.Email,
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

func (c *authClient) GoogleLogin(ctx context.Context, request models.GoogleLoginRequest, retryConfig models.RetryConfig) (*auth.GoogleLoginResponse, error) {

	var res *auth.GoogleLoginResponse
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
			res, err = c.Server.GoogleLogin(ctx1, &auth.GoogleLoginRequest{
				Token: request.Token,
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

func (c *authClient) ChangeUserName(ctx context.Context, request models.ChangeUserNameRequest, retryConfig models.RetryConfig) (*auth.ChangeUserNameResponse, error) {
	// Retrieve the "userId" from the context.
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	var res *auth.ChangeUserNameResponse
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
			res, err = c.Server.ChangeUserName(ctx1, &auth.ChangeUserNameRequest{
				UserId:   userId,
				UserName: request.UserName,
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
func (c *authClient) ChangeEmail(ctx context.Context, request models.ChangeEmailRequest, retryConfig models.RetryConfig) (*auth.ChangeEmailResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	var res *auth.ChangeEmailResponse
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
			res, err = c.Server.ChangeEmail(ctx1, &auth.ChangeEmailRequest{
				UserId: userId,
				Email:  request.Email,
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

func (c *authClient) ChangePassword(ctx context.Context, request models.ChangePasswordRequest, retryConfig models.RetryConfig) (*auth.ChangePasswordResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	var res *auth.ChangePasswordResponse
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
			res, err = c.Server.ChangePassword(ctx1, &auth.ChangePasswordRequest{
				UserId:   userId,
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
func (c *authClient) ChangeEmailVerifyOtp(ctx context.Context, request models.ChangeEmailVerifyOtpRequest, retryConfig models.RetryConfig) (*auth.ChangeEmailVerifyOtpResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	var res *auth.ChangeEmailVerifyOtpResponse
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
			res, err = c.Server.ChangeEmailVerifyOtp(ctx1, &auth.ChangeEmailVerifyOtpRequest{
				UserId: userId,
				Email:  request.Email,
				Otp:    request.Otp,
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

func (c *authClient) ChangePhoneNumberOtp(ctx context.Context, request models.ChangePhoneNumberOtpRequest, retryConfig models.RetryConfig) (*auth.ChangePhoneNumberOtpResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	var res *auth.ChangePhoneNumberOtpResponse
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
			res, err = c.Server.ChangePhoneNumberOtp(ctx1, &auth.ChangePhoneNumberOtpRequest{
				UserId:      userId,
				PhoneNumber: request.PhoneNumber,
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

func (c *authClient) ChangePhoneNumber(ctx context.Context, request models.ChangePhoneNumberRequest, retryConfig models.RetryConfig) (*auth.ChangePhoneNumberResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	var res *auth.ChangePhoneNumberResponse
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
			res, err = c.Server.ChangePhoneNumber(ctx1, &auth.ChangePhoneNumberRequest{
				UserId:      userId,
				PhoneNumber: request.PhoneNumber,
				Otp:         request.Otp,
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

func (c *authClient) ChangeAvatar(ctx context.Context, request models.ChangeAvatarRequest, retryConfig models.RetryConfig) (*auth.ChangeAvatarResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	var res *auth.ChangeAvatarResponse
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
			res, err = c.Server.ChangeAvatar(ctx1, &auth.ChangeAvatarRequest{
				UserId:   userId,
				AvatarId: request.AvatarId,
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

func (c *authClient) RemoveAvatar(ctx context.Context, retryConfig models.RetryConfig) (*auth.RemoveAvatarResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	var res *auth.RemoveAvatarResponse
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
			res, err = c.Server.RemoveAvatar(ctx1, &auth.RemoveAvatarRequest{
				UserId: userId,
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

func (c *authClient) CreateCommunity(ctx context.Context, request models.CreateCommunityRequest) (*auth.CreateCommunityResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := c.Server.CreateCommunity(ctx, &auth.CreateCommunityRequest{
		AdminId:        userId,
		CommunityName:  request.CommunityName,
		Members:        request.Members,
		Description:    request.Description,
		JoinedType:     request.JoinedType,
		CommunityImage: request.CommunityImage,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) JoinCommunity(ctx context.Context, request models.JoinCommunityRequest) (*auth.JoinCommunityResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	log.Printf("Calling JoinCommunity with UserId: %s, CommunityId: %s, Message: %s", userId, request.CommunityId, request.Message)
	res, err := c.Server.JoinCommunity(ctx, &auth.JoinCommunityRequest{
		UserId:      userId,
		CommunityId: request.CommunityId,
		Message:     request.Message,
	})

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) LeaveCommunity(ctx context.Context, request models.LeaveCommunityRequest) (*auth.LeaveCommunityResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := c.Server.LeaveCommunity(ctx, &auth.LeaveCommunityRequest{
		UserId:      userId,
		CommunityId: request.CommunityId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (c *authClient) AcceptJoinCommunity(ctx context.Context, request models.AcceptJoinCommunityRequest) (*auth.AcceptJoinCommunityResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := c.Server.AcceptJoinCommunity(ctx, &auth.AcceptJoinCommunityRequest{
		UserId:      request.RequstedUserId,
		CommunityId: request.CommunityId,
		AdminId:     userId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) RemoveMember(ctx context.Context, request models.RemoveMemberRequest) (*auth.RemoveMemberResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := c.Server.RemoveMember(ctx, &auth.RemoveMemberRequest{
		CommunityId: request.CommunityId,
		AdminId:     userId,
		UserId:      request.MemberUserId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) AddModerator(ctx context.Context, request models.AddModeratorRequest) (*auth.AddModeratorResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := c.Server.AddModerator(ctx, &auth.AddModeratorRequest{
		AdminId:     userId,
		UserId:      request.ModeratorUserId,
		CommunityId: request.CommunityId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) AddMember(ctx context.Context, request models.AddMemberRequest) (*auth.AddMemberResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := c.Server.AddMember(ctx, &auth.AddMemberRequest{
		UserId:      request.MemberUserId,
		AdminId:     userId,
		CommunityId: request.CommunityId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) ChangeCommunityJoinType(ctx context.Context, request models.ChangeCommunityJoinTypeRequest) (*auth.ChangeCommunityJoinTypeResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := c.Server.ChangeCommunityJoinType(ctx, &auth.ChangeCommunityJoinTypeRequest{
		UserId:      userId,
		CommunityId: request.CommunityId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) DeleteCommunity(ctx context.Context, request models.DeleteCommunityRequest) (*auth.DeleteCommunityResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := c.Server.DeleteCommunity(ctx, &auth.DeleteCommunityRequest{
		UserId:      userId,
		CommunityId: request.CommunityId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) GetInterstsUser(ctx context.Context) (*auth.GetInterstsUserResponse, error) {
	res, err := c.Server.GetInterest(ctx, &auth.GetInterestRequest{})
	if err != nil {
		return nil, err
	}

	userResponse := &auth.GetInterstsUserResponse{
		Status:    res.Status,
		Interests: make([]*auth.UserInterest, len(res.Interests)),
		Message:   res.Message,
		Error:     res.Error,
	}

	for i, interest := range res.Interests {
		userResponse.Interests[i] = &auth.UserInterest{
			Interest: interest.Interest,
		}
	}

	return userResponse, nil
}

func (c *authClient) GetUserByName(ctx context.Context, request models.GetUserByNameRequest) (*auth.GetUserByNameResponse, error) {
	res, err := c.Server.GetUserByName(ctx, &auth.GetUserByNameRequest{
		UserName: request.UserName,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) GetActiveCommunity(ctx context.Context) (*auth.GetActiveCommunityResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}
	res, err := c.Server.GetActiveCommunity(ctx, &auth.GetActiveCommunityRequest{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (c *authClient) GetCommunityById(ctx context.Context, request models.GetCommunityByIdRequest) (*auth.GetCommunityByIdResponse, error) {
	res, err := c.Server.GetCommunityById(ctx, &auth.GetCommunityByIdRequest{
		CommunityId: request.CommunityId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) ValidateCommunityName(ctx context.Context, request models.ValidateCommunityNameRequest) (*auth.ValidateCommunityNameResponse, error) {
	res, err := c.Server.ValidateCommunityName(ctx, &auth.ValidateCommunityNameRequest{
		CommunityName: request.CommunityName,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) GetAllCommunity(ctx context.Context) (*auth.GetAllCommunityResponse, error) {
	res, err := c.Server.GetAllCommunity(ctx, &auth.GetAllCommunityRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) GetUserDetails(ctx context.Context) (*auth.GetUserDetailsResponse, error) {

	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := c.Server.GetUserDetails(ctx, &auth.GetUserDetailsRequest{
		UserID: userId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (c *authClient) GetJoinedCommunity(ctx context.Context) (*auth.GetJoinedCommunityResponse, error) {

	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := c.Server.GetJoinedCommunity(ctx, &auth.GetJoinedCommunityRequest{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (c *authClient) SearchCommunity(ctx context.Context, request models.SearchCommunityRequest) (*auth.SearchCommunityResponse, error) {

	res, err := c.Server.SearchCommunity(ctx, &auth.SearchCommunityRequest{
		CommunityName: request.CommunityName,
	})
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (c *authClient) ReportVideo(ctx context.Context, request models.ReportVideoRequest) (*auth.ReportVideoResponse, error) {
	res, err := c.Server.ReportVideo(ctx, &auth.ReportVideoRequest{
		VideoId:      request.VideoId,
		Reason:       request.Reason,
		ReportedUser: request.ReportedUser,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
