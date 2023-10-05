package interfaces

import (
	"context"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb/monit"
)

type MonitaizationClient interface {
	HealthCheck(context.Context, models.Request) (*monit.Response, error)
	CreateWallet(context.Context, models.CreateWalletRequest) (*monit.CreateWalletResponse, error)
	UpdateWallet(context.Context, models.UpdateWalletRequest) (*monit.UpdateWalletResponse, error)
	ParticipationReward(context.Context, models.ParticipationRewardRequest) (*monit.ParticipationRewardResponse, error)
	UserRewardHistory(context.Context, models.UserRewardHistoryRequest) (*monit.UserRewardHistoryResponse, error)
	GetWallet(context.Context) (*monit.GetWalletResponse, error)
	ExclusiveContent(context.Context, string) (*monit.ExclusiveContentResponse, error)
}
