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

func (c *authClient) UserSignup(ctx context.Context, request models.RegisterRequestBody, retryConfig models.RetryConfig) (*pb.SignupResponse, error) {

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

func (c *authClient) OtpRequest(ctx context.Context, request models.OtpValidation, retryConfig models.RetryConfig) (*pb.OtpSignUpResponse, error) {
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

func (c *authClient) ValidName(ctx context.Context, request models.ValidName, retryConfig models.RetryConfig) (*pb.ValidNameResponse, error) {
	var res *pb.ValidNameResponse
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
			res, err = c.Server.ValidName(ctx1, &pb.ValidNameRequest{
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

func (c *authClient) ResendOtp(ctx context.Context, request models.ResendOtp, retryConfig models.RetryConfig) (*pb.ResendOtpResponse, error) {

	var res *pb.ResendOtpResponse
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
			res, err = c.Server.ResendOtp(ctx1, &pb.ResendOtpRequest{
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

func (c *authClient) ForgotPasswordOtp(ctx context.Context, request models.ForgotPasswordOtpRequest, retryConfig models.RetryConfig) (*pb.ForgotPasswordOtpResponse, error) {

	var res *pb.ForgotPasswordOtpResponse
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
			res, err = c.Server.ForgotPasswordOtp(ctx1, &pb.ForgotPasswordOtpRequest{
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

func (c *authClient) ForgotPasswordValidateOtp(ctx context.Context, request models.ForgotPasswordValidateOtpRequest, retryConfig models.RetryConfig) (*pb.ForgotPasswordValidateOtpResponse, error) {

	var res *pb.ForgotPasswordValidateOtpResponse
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
			res, err = c.Server.ForgotPasswordValidateOtp(ctx1, &pb.ForgotPasswordValidateOtpRequest{
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

func (c *authClient) ForgotPasswordChangePassword(ctx context.Context, request models.ForgotPasswordChangePasswordRequest, retryConfig models.RetryConfig) (*pb.ForgotPasswordChangePasswordResponse, error) {

	var res *pb.ForgotPasswordChangePasswordResponse
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
			res, err = c.Server.ForgotPasswordChangePassword(ctx1, &pb.ForgotPasswordChangePasswordRequest{
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

func (c *authClient) ValidateUser(ctx context.Context, request models.ValidateUserRequest, retryConfig models.RetryConfig) (*pb.ValidateUserResponse, error) {

	var res *pb.ValidateUserResponse
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
			res, err = c.Server.ValidateUser(ctx1, &pb.ValidateUserRequest{
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

func (c *authClient) GoogleLogin(ctx context.Context, request models.GoogleLoginRequest, retryConfig models.RetryConfig) (*pb.GoogleLoginResponse, error) {

	var res *pb.GoogleLoginResponse
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
			res, err = c.Server.GoogleLogin(ctx1, &pb.GoogleLoginRequest{
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

func (c *authClient) ChangeUserName(ctx context.Context, request models.ChangeUserNameRequest, retryConfig models.RetryConfig) (*pb.ChangeUserNameResponse, error) {
	// Retrieve the "userId" from the context.
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	var res *pb.ChangeUserNameResponse
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
			res, err = c.Server.ChangeUserName(ctx1, &pb.ChangeUserNameRequest{
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
func (c *authClient) ChangeEmail(ctx context.Context, request models.ChangeEmailRequest, retryConfig models.RetryConfig) (*pb.ChangeEmailResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	var res *pb.ChangeEmailResponse
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
			res, err = c.Server.ChangeEmail(ctx1, &pb.ChangeEmailRequest{
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

func (c *authClient) ChangePassword(ctx context.Context, request models.ChangePasswordRequest, retryConfig models.RetryConfig) (*pb.ChangePasswordResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	var res *pb.ChangePasswordResponse
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
			res, err = c.Server.ChangePassword(ctx1, &pb.ChangePasswordRequest{
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
func (c *authClient) ChangeEmailVerifyOtp(ctx context.Context, request models.ChangeEmailVerifyOtpRequest, retryConfig models.RetryConfig) (*pb.ChangeEmailVerifyOtpResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	var res *pb.ChangeEmailVerifyOtpResponse
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
			res, err = c.Server.ChangeEmailVerifyOtp(ctx1, &pb.ChangeEmailVerifyOtpRequest{
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

func (c *authClient) ChangePhoneNumberOtp(ctx context.Context, request models.ChangePhoneNumberOtpRequest, retryConfig models.RetryConfig) (*pb.ChangePhoneNumberOtpResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	var res *pb.ChangePhoneNumberOtpResponse
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
			res, err = c.Server.ChangePhoneNumberOtp(ctx1, &pb.ChangePhoneNumberOtpRequest{
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

func (c *authClient) ChangePhoneNumber(ctx context.Context, request models.ChangePhoneNumberRequest, retryConfig models.RetryConfig) (*pb.ChangePhoneNumberResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	var res *pb.ChangePhoneNumberResponse
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
			res, err = c.Server.ChangePhoneNumber(ctx1, &pb.ChangePhoneNumberRequest{
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

func (c *authClient) ChangeAvatar(ctx context.Context, request models.ChangeAvatarRequest, retryConfig models.RetryConfig) (*pb.ChangeAvatarResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	var res *pb.ChangeAvatarResponse
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
			res, err = c.Server.ChangeAvatar(ctx1, &pb.ChangeAvatarRequest{
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

func (c *authClient) RemoveAvatar(ctx context.Context, retryConfig models.RetryConfig) (*pb.RemoveAvatarResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	var res *pb.RemoveAvatarResponse
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
			res, err = c.Server.RemoveAvatar(ctx1, &pb.RemoveAvatarRequest{
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

func (c *authClient) CreateCommunity(ctx context.Context, request models.CreateCommunityRequest) (*pb.CreateCommunityResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := c.Server.CreateCommunity(ctx, &pb.CreateCommunityRequest{
		AdminId:       userId,
		CommunityName: request.CommunityName,
		Moderators:    request.Moderator,
		Members:       request.Members,
		Description:   request.Description,
		JoinedType:    request.JoinedType,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) JoinCommunity(ctx context.Context, request models.JoinCommunityRequest) (*pb.JoinCommunityResponse, error) {

	res, err := c.Server.JoinCommunity(ctx, &pb.JoinCommunityRequest{
		UserId:      request.MemberUserID,
		CommunityId: request.CommunityId,
		JoinType:    request.JoinType,
		Message:     request.Message,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) LeaveCommunity(ctx context.Context, request models.LeaveCommunityRequest) (*pb.LeaveCommunityResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := c.Server.LeaveCommunity(ctx, &pb.LeaveCommunityRequest{
		UserId:      userId,
		CommunityId: request.CommunityId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (c *authClient) AcceptJoinCommunity(ctx context.Context, request models.AcceptJoinCommunityRequest) (*pb.AcceptJoinCommunityResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := c.Server.AcceptJoinCommunity(ctx, &pb.AcceptJoinCommunityRequest{
		UserId:      request.RequstedUserId,
		CommunityId: request.CommunityId,
		AdminId:     userId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) RemoveMember(ctx context.Context, request models.RemoveMemberRequest) (*pb.RemoveMemberResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := c.Server.RemoveMember(ctx, &pb.RemoveMemberRequest{
		CommunityId: request.CommunityId,
		AdminId:     userId,
		UserId:      request.MemberUserId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) AddModerator(ctx context.Context, request models.AddModeratorRequest) (*pb.AddModeratorResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := c.Server.AddModerator(ctx, &pb.AddModeratorRequest{
		AdminId:     userId,
		UserId:      request.ModeratorUserId,
		CommunityId: request.CommunityId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) AddMember(ctx context.Context, request models.AddMemberRequest) (*pb.AddMemberResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := c.Server.AddMember(ctx, &pb.AddMemberRequest{
		UserId:      request.MemberUserId,
		AdminId:     userId,
		CommunityId: request.CommunityId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) ChangeCommunityJoinType(ctx context.Context, request models.ChangeCommunityJoinTypeRequest) (*pb.ChangeCommunityJoinTypeResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := c.Server.ChangeCommunityJoinType(ctx, &pb.ChangeCommunityJoinTypeRequest{
		UserId:      userId,
		CommunityId: request.CommunityId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *authClient) DeleteCommunity(ctx context.Context, request models.DeleteCommunityRequest) (*pb.DeleteCommunityResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := c.Server.DeleteCommunity(ctx, &pb.DeleteCommunityRequest{
		UserId:      userId,
		CommunityId: request.CommunityId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
