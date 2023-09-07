package interfaces

import (
	"context"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb"
)

type AuthClient interface {
	UserSignup(context.Context, models.RegisterRequestBody, models.RetryConfig) (*pb.SignupResponse, error)
	OtpRequest(context.Context, models.OtpValidation, models.RetryConfig) (*pb.OtpSignUpResponse, error)
	UserLogin(context.Context, models.LoginRequestBody, models.RetryConfig) (*pb.LoginResponse, error)
	ValidName(context.Context, models.ValidName, models.RetryConfig) (*pb.ValidNameResponse, error)
	ResendOtp(context.Context, models.ResendOtp, models.RetryConfig) (*pb.ResendOtpResponse, error)
	ForgotPasswordOtp(context.Context, models.ForgotPasswordOtpRequest, models.RetryConfig) (*pb.ForgotPasswordOtpResponse, error)
	ForgotPasswordValidateOtp(context.Context, models.ForgotPasswordValidateOtpRequest, models.RetryConfig) (*pb.ForgotPasswordValidateOtpResponse, error)
	ForgotPasswordChangePassword(context.Context, models.ForgotPasswordChangePasswordRequest, models.RetryConfig) (*pb.ForgotPasswordChangePasswordResponse, error)
	ValidateUser(context.Context, models.ValidateUserRequest, models.RetryConfig) (*pb.ValidateUserResponse, error)
	GoogleLogin(context.Context, models.GoogleLoginRequest, models.RetryConfig) (*pb.GoogleLoginResponse, error)
	ChangeUserName(context.Context, models.ChangeUserNameRequest, models.RetryConfig) (*pb.ChangeUserNameResponse, error)
	ChangeEmail(context.Context, models.ChangeEmailRequest, models.RetryConfig) (*pb.ChangeEmailResponse, error)
	ChangePassword(context.Context, models.ChangePasswordRequest, models.RetryConfig) (*pb.ChangePasswordResponse, error)
	ChangeEmailVerifyOtp(context.Context, models.ChangeEmailVerifyOtpRequest, models.RetryConfig) (*pb.ChangeEmailVerifyOtpResponse, error)
	ChangePhoneNumberOtp(context.Context, models.ChangePhoneNumberOtpRequest, models.RetryConfig) (*pb.ChangePhoneNumberOtpResponse, error)
	ChangePhoneNumber(context.Context, models.ChangePhoneNumberRequest, models.RetryConfig) (*pb.ChangePhoneNumberResponse, error)
	ChangeAvatar(context.Context, models.ChangeAvatarRequest, models.RetryConfig) (*pb.ChangeAvatarResponse, error)
	RemoveAvatar(context.Context, models.RetryConfig) (*pb.RemoveAvatarResponse, error)
	CreateCommunity(context.Context, models.CreateCommunityRequest) (*pb.CreateCommunityResponse, error)
	JoinCommunity(context.Context, models.JoinCommunityRequest) (*pb.JoinCommunityResponse, error)
	LeaveCommunity(context.Context, models.LeaveCommunityRequest) (*pb.LeaveCommunityResponse, error)
	AcceptJoinCommunity(context.Context, models.AcceptJoinCommunityRequest) (*pb.AcceptJoinCommunityResponse, error)
	RemoveMember(context.Context, models.RemoveMemberRequest) (*pb.RemoveMemberResponse, error)
	AddModerator(context.Context, models.AddModeratorRequest) (*pb.AddModeratorResponse, error)
	AddMember(context.Context, models.AddMemberRequest) (*pb.AddMemberResponse, error)
	ChangeCommunityJoinType(context.Context, models.ChangeCommunityJoinTypeRequest) (*pb.ChangeCommunityJoinTypeResponse, error)
	DeleteCommunity(context.Context, models.DeleteCommunityRequest) (*pb.DeleteCommunityResponse, error)
}
