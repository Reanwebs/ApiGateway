package client

import (
	"context"
	"errors"
	"fmt"
	"gateway/pkg/common/client/interfaces"
	"gateway/pkg/common/config"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type monitaizationClient struct {
	Server pb.MonitizationClient
}

func InitMonitizationClient(c *config.Config) (interfaces.MonitaizationClient, error) {
	cc, err := grpc.Dial(c.MonitaizationService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return NewMonitizationClient(pb.NewMonitizationClient(cc)), nil
}

func NewMonitizationClient(server pb.MonitizationClient) interfaces.MonitaizationClient {
	return &monitaizationClient{
		Server: server,
	}
}

func (m *monitaizationClient) HealthCheck(ctx context.Context, request models.Request) (*pb.Response, error) {

	res, err := m.Server.HealthCheck(ctx, &pb.Request{
		Data: request.Data,
	})
	if err != nil {
		return nil, err
	}

	return res, err
}

func (m *monitaizationClient) CreateWallet(ctx context.Context, request models.CreateWalletRequest) (*pb.CreateWalletResponse, error) {
	res, err := m.Server.CreateWallet(ctx, &pb.CreateWalletRequest{
		UserID: request.UserID,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *monitaizationClient) UpdateWallet(ctx context.Context, request models.UpdateWalletRequest) (*pb.UpdateWalletResponse, error) {
	res, err := m.Server.UpdateWallet(ctx, &pb.UpdateWalletRequest{
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

func (m *monitaizationClient) ParticipationReward(ctx context.Context, request models.ParticipationRewardRequest) (*pb.ParticipationRewardResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}
	res, err := m.Server.ParticipationReward(ctx, &pb.ParticipationRewardRequest{
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

func (m *monitaizationClient) UserRewardHistory(ctx context.Context, request models.UserRewardHistoryRequest) (*pb.UserRewardHistoryResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := m.Server.UserRewardHistory(ctx, &pb.UserRewardHistoryRequest{
		UserID: userId,
		Sort:   request.Sort,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *monitaizationClient) GetWallet(ctx context.Context) (*pb.GetWalletResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}
	res, err := m.Server.GetWallet(ctx, &pb.GetWalletRequest{
		UserID: userId,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
