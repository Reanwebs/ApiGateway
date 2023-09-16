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
	ManageInterest(context.Context, models.ManageInterestRequest) (*pb.ManageInterestResponse, error)
	ManageCommunity(context.Context, models.ManageCommunityRequest) (*pb.ManageCommunityResponse, error)
	GetAllCommunity(context.Context) (*pb.GetAllCommunityResponse, error)
}
