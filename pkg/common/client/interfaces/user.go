package interfaces

import (
	"context"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb/auth"
)

type AuthClient interface {
	UserSignup(context.Context, models.RegisterRequestBody, models.RetryConfig) (*auth.SignupResponse, error)
	OtpRequest(context.Context, models.OtpValidation, models.RetryConfig) (*auth.OtpSignUpResponse, error)
	UserLogin(context.Context, models.LoginRequestBody, models.RetryConfig) (*auth.LoginResponse, error)
	ValidName(context.Context, models.ValidName, models.RetryConfig) (*auth.ValidNameResponse, error)
	ResendOtp(context.Context, models.ResendOtp, models.RetryConfig) (*auth.ResendOtpResponse, error)
	ForgotPasswordOtp(context.Context, models.ForgotPasswordOtpRequest, models.RetryConfig) (*auth.ForgotPasswordOtpResponse, error)
	ForgotPasswordValidateOtp(context.Context, models.ForgotPasswordValidateOtpRequest, models.RetryConfig) (*auth.ForgotPasswordValidateOtpResponse, error)
	ForgotPasswordChangePassword(context.Context, models.ForgotPasswordChangePasswordRequest, models.RetryConfig) (*auth.ForgotPasswordChangePasswordResponse, error)
	ValidateUser(context.Context, models.ValidateUserRequest, models.RetryConfig) (*auth.ValidateUserResponse, error)
	GoogleLogin(context.Context, models.GoogleLoginRequest, models.RetryConfig) (*auth.GoogleLoginResponse, error)
	ChangeUserName(context.Context, models.ChangeUserNameRequest, models.RetryConfig) (*auth.ChangeUserNameResponse, error)
	ChangeEmail(context.Context, models.ChangeEmailRequest, models.RetryConfig) (*auth.ChangeEmailResponse, error)
	ChangePassword(context.Context, models.ChangePasswordRequest, models.RetryConfig) (*auth.ChangePasswordResponse, error)
	ChangeEmailVerifyOtp(context.Context, models.ChangeEmailVerifyOtpRequest, models.RetryConfig) (*auth.ChangeEmailVerifyOtpResponse, error)
	ChangePhoneNumberOtp(context.Context, models.ChangePhoneNumberOtpRequest, models.RetryConfig) (*auth.ChangePhoneNumberOtpResponse, error)
	ChangePhoneNumber(context.Context, models.ChangePhoneNumberRequest, models.RetryConfig) (*auth.ChangePhoneNumberResponse, error)
	ChangeAvatar(context.Context, models.ChangeAvatarRequest, models.RetryConfig) (*auth.ChangeAvatarResponse, error)
	RemoveAvatar(context.Context, models.RetryConfig) (*auth.RemoveAvatarResponse, error)
	CreateCommunity(context.Context, models.CreateCommunityRequest) (*auth.CreateCommunityResponse, error)
	JoinCommunity(context.Context, models.JoinCommunityRequest) (*auth.JoinCommunityResponse, error)
	LeaveCommunity(context.Context, models.LeaveCommunityRequest) (*auth.LeaveCommunityResponse, error)
	AcceptJoinCommunity(context.Context, models.AcceptJoinCommunityRequest) (*auth.AcceptJoinCommunityResponse, error)
	RemoveMember(context.Context, models.RemoveMemberRequest) (*auth.RemoveMemberResponse, error)
	AddModerator(context.Context, models.AddModeratorRequest) (*auth.AddModeratorResponse, error)
	AddMember(context.Context, models.AddMemberRequest) (*auth.AddMemberResponse, error)
	ChangeCommunityJoinType(context.Context, models.ChangeCommunityJoinTypeRequest) (*auth.ChangeCommunityJoinTypeResponse, error)
	DeleteCommunity(context.Context, models.DeleteCommunityRequest) (*auth.DeleteCommunityResponse, error)
	GetInterstsUser(context.Context) (*auth.GetInterstsUserResponse, error)
	GetUserByName(context.Context, models.GetUserByNameRequest) (*auth.GetUserByNameResponse, error)
	GetActiveCommunity(context.Context) (*auth.GetActiveCommunityResponse, error)
	GetCommunityById(context.Context, models.GetCommunityByIdRequest) (*auth.GetCommunityByIdResponse, error)
	ValidateCommunityName(context.Context, models.ValidateCommunityNameRequest) (*auth.ValidateCommunityNameResponse, error)
	GetJoinedCommunity(context.Context) (*auth.GetJoinedCommunityResponse, error)
	GetUserDetails(context.Context) (*auth.GetUserDetailsResponse, error)
	SearchCommunity(context.Context, models.SearchCommunityRequest) (*auth.SearchCommunityResponse, error)
	ReportVideo(ctx context.Context, request models.ReportVideoRequest) (*auth.ReportVideoResponse, error)
}
