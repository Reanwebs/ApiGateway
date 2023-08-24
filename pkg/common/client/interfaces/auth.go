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
}
