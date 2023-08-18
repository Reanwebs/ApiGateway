package interfaces

import (
	"context"
	"gateway/pkg/auth/common/models"
	"gateway/pkg/auth/common/pb"
)

type AuthClient interface {
	UserSignup(context.Context, models.RegisterRequestBody) (*pb.SignupResponse, error)
	OtpValidation(context.Context, models.OtpValidation) (*pb.OtpValidationResponse, error)
	UserLogin(context.Context, models.LoginRequestBody) (*pb.LoginResponse, error)
}
