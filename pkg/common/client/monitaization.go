package client

import (
	"context"
	"errors"
	"gateway/pkg/common/client/interfaces"
	"gateway/pkg/common/config"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb/monit"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type monitaizationClient struct {
	Server monit.MonitizationClient
}

func InitMonitizationClient(c *config.Config) (interfaces.MonitaizationClient, error) {
	cc, err := grpc.Dial(c.MonitaizationService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return NewMonitizationClient(monit.NewMonitizationClient(cc)), nil
}

func NewMonitizationClient(server monit.MonitizationClient) interfaces.MonitaizationClient {
	return &monitaizationClient{
		Server: server,
	}
}

func (m *monitaizationClient) HealthCheck(ctx context.Context, request models.Request) (*monit.Response, error) {

	res, err := m.Server.HealthCheck(ctx, &monit.Request{
		Data: request.Data,
	})
	if err != nil {
		return nil, err
	}

	return res, err
}

func (m *monitaizationClient) CreateWallet(ctx context.Context, request models.CreateWalletRequest) (*monit.CreateWalletResponse, error) {
	res, err := m.Server.CreateWallet(ctx, &monit.CreateWalletRequest{
		UserID: request.UserID,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *monitaizationClient) UpdateWallet(ctx context.Context, request models.UpdateWalletRequest) (*monit.UpdateWalletResponse, error) {
	res, err := m.Server.UpdateWallet(ctx, &monit.UpdateWalletRequest{
		UserID:   request.UserID,
		Type:     request.Type,
		Reason:   request.Reason,
		UserName: request.UserName,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *monitaizationClient) ParticipationReward(ctx context.Context, request models.ParticipationRewardRequest) (*monit.ParticipationRewardResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		log.Println("userId not found in context.")
		return nil, errors.New("login again")
	}
	res, err := m.Server.ParticipationReward(ctx, &monit.ParticipationRewardRequest{
		UserID:         userId,
		ConferenceID:   request.ConferenceID,
		ConferenceType: request.ConferenceType,
		Minutes:        request.Minutes,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *monitaizationClient) UserRewardHistory(ctx context.Context, request models.UserRewardHistoryRequest) (*monit.UserRewardHistoryResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		log.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := m.Server.UserRewardHistory(ctx, &monit.UserRewardHistoryRequest{
		UserID: userId,
		Sort:   request.Sort,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *monitaizationClient) GetWallet(ctx context.Context) (*monit.GetWalletResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		log.Println("userId not found in context.")
		return nil, errors.New("login again")
	}
	res, err := m.Server.GetWallet(ctx, &monit.GetWalletRequest{
		UserID: userId,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *monitaizationClient) ExclusiveContent(ctx context.Context, videoId string) (*monit.ExclusiveContentResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		log.Println("userId not found in context.")
		return nil, errors.New("login again")
	}
	res, err := m.Server.ExclusiveContent(ctx, &monit.ExclusiveContentRequest{
		UserID:  userId,
		VideoID: videoId,
	})
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (m *monitaizationClient) EmailNotification(ctx context.Context, request models.EmailNotificationRequest) (*monit.EmailNotificationResponse, error) {
	var monitEmails []*monit.Emails

	for _, email := range request.Emails {
		monitEmail := &monit.Emails{
			Email: email,
		}
		monitEmails = append(monitEmails, monitEmail)
	}

	res, err := m.Server.EmailNotification(ctx, &monit.EmailNotificationRequest{
		To:      monitEmails,
		Subject: request.Subject,
		Body:    request.Body,
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
