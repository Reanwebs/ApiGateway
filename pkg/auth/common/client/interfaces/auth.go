package interfaces

import (
	"context"
	"gateway/pkg/auth/common/pb"
	"mime/multipart"
)

type AuthClient interface {
	UserSignup(context.Context, *multipart.FileHeader) (*pb.SignupResponse, error)
	UserLogin(context.Context, string, string) (pb.LoginResponse, error)
}
