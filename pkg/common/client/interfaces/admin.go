package interfaces

import (
	"context"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb"
)

type AdminClient interface {
	AdminLogin(context.Context, models.AdminLoginRequest) (*pb.AdminLoginResponse, error)
	GetUsers(context.Context) (*pb.GetUsersResponse, error)
	ManageUser(context.Context, models.ManageUserRequest) (*pb.ManageUserResponse, error)
	GetInterest(context.Context) (*pb.GetInterestResponse, error)
	AddInterest(context.Context, models.AddInterestRequest) (*pb.AddInterestResponse, error)
}
