package interfaces

import (
	"context"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb"
)

type MonitaizationClient interface {
	HealthCheck(context.Context, models.Request) (*pb.Response, error)
	CreateWallet(context.Context, models.CreateWalletRequest) (*pb.CreateWalletResponse, error)
	UpdateWallet(context.Context, models.UpdateWalletRequest) (*pb.UpdateWalletResponse, error)
	ParticipationReward(context.Context, models.ParticipationRewardRequest) (*pb.ParticipationRewardResponse, error)
	UserRewardHistory(context.Context, models.UserRewardHistoryRequest) (*pb.UserRewardHistoryResponse, error)
	GetWallet(context.Context) (*pb.GetWalletResponse, error)
}
