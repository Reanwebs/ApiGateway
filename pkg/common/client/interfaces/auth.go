package interfaces

import (
	"context"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb"
)

type AuthClient interface {
	UserSignup(context.Context, models.RegisterRequestBody) (*pb.SignupResponse, error)
	OtpRequest(context.Context, models.OtpValidation) (*pb.OtpSignUpResponse, error)
	UserLogin(context.Context, models.LoginRequestBody) (*pb.LoginResponse, error)
	ValidName(context.Context, models.ValidName) (*pb.ValidNameResponse, error)
	ResendOtp(context.Context, models.ResendOtp) (*pb.ResendOtpResponse, error)
	ForgotPasswordOtp(context.Context, models.ForgotPasswordOtpRequest) (*pb.ForgotPasswordOtpResponse, error)
	ForgotPasswordValidateOtp(context.Context, models.ForgotPasswordValidateOtpRequest) (*pb.ForgotPasswordValidateOtpResponse, error)
	ForgotPasswordChangePassword(context.Context, models.ForgotPasswordChangePasswordRequest) (*pb.ForgotPasswordChangePasswordResponse, error)
}
