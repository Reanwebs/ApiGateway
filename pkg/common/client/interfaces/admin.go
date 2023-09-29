package interfaces

import (
	"context"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb/auth"
)

type AdminClient interface {
	AdminLogin(context.Context, models.AdminLoginRequest) (*auth.AdminLoginResponse, error)
	GetUsers(context.Context) (*auth.GetUsersResponse, error)
	ManageUser(context.Context, models.ManageUserRequest) (*auth.ManageUserResponse, error)
	GetInterest(context.Context) (*auth.GetInterestResponse, error)
	AddInterest(context.Context, models.AddInterestRequest) (*auth.AddInterestResponse, error)
	ManageInterest(context.Context, models.ManageInterestRequest) (*auth.ManageInterestResponse, error)
	ManageCommunity(context.Context, models.ManageCommunityRequest) (*auth.ManageCommunityResponse, error)
	GetAllCommunity(context.Context) (*auth.GetAllCommunityResponse, error)
}
