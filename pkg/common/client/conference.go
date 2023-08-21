package client

import (
	"context"
	"gateway/pkg/common/client/interfaces"
	"gateway/pkg/common/config"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type conferenceClient struct {
	Server pb.ConferenceClient
}

func InitConferenceClient(c *config.Config) (interfaces.ConferenceClient, error) {
	cc, err := grpc.Dial(c.ConferenceService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return NewConferenceClient(pb.NewConferenceClient(cc)), nil
}

func NewConferenceClient(server pb.ConferenceClient) interfaces.ConferenceClient {
	return &conferenceClient{
		Server: server,
	}
}

func (c *conferenceClient) HealthCheck(ctx context.Context, request string) (*pb.Response, error) {
	res, err := c.Server.HealthCheck(ctx, &pb.Request{
		Data: request,
	})
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (c *conferenceClient) StartConference(ctx context.Context, request models.StartConferenceRequest) (*pb.StartConferenceResponse, error) {
	res, err := c.Server.StartConference(ctx, &pb.StartConferenceRequest{
		UserID:           request.UserID,
		Type:             request.Type,
		Title:            request.Title,
		Description:      request.Description,
		Interest:         request.Interest,
		Recording:        request.Recording,
		Chat:             request.Chat,
		Broadcast:        request.Broadcast,
		Participantlimit: request.Participantlimit,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) JoinConference(ctx context.Context, request models.JoinConferenceRequest) (*pb.JoinConferenceResponse, error) {
	res, err := c.Server.JoinConference(ctx, &pb.JoinConferenceRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) AcceptJoining(ctx context.Context, request models.AcceptJoiningRequest) (*pb.AcceptJoiningResponse, error) {
	res, err := c.Server.AcceptJoining(ctx, &pb.AcceptJoiningRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) DeclineJoining(ctx context.Context, request models.DeclineJoiningRequest) (*pb.DeclineJoiningResponse, error) {
	res, err := c.Server.DeclineJoining(ctx, &pb.DeclineJoiningRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) LeaveConference(ctx context.Context, request models.LeaveConferenceRequest) (*pb.LeaveConferenceResponse, error) {
	res, err := c.Server.LeaveConference(ctx, &pb.LeaveConferenceRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) RemoveParticipant(ctx context.Context, request models.RemoveParticipantRequest) (*pb.RemoveParticipantResponse, error) {
	res, err := c.Server.RemoveParticipant(ctx, &pb.RemoveParticipantRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
		Block:        request.Block,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) ToggleCamera(ctx context.Context, request models.ToggleCameraRequest) (*pb.ToggleCameraResponse, error) {
	res, err := c.Server.ToggleCamera(ctx, &pb.ToggleCameraRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) ToggleMic(ctx context.Context, request models.ToggleMicRequest) (*pb.ToggleMicResponse, error) {
	res, err := c.Server.ToggleMic(ctx, &pb.ToggleMicRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) ToggleParticipantCamera(ctx context.Context, request models.ToggleParticipantCameraRequest) (*pb.ToggleParticipantCameraResponse, error) {
	res, err := c.Server.ToggleParticipantCamera(ctx, &pb.ToggleParticipantCameraRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) ToggleParticipantMic(ctx context.Context, request models.ToggleParticipantMicRequest) (*pb.ToggleParticipantMicResponse, error) {
	res, err := c.Server.ToggleParticipantMic(ctx, &pb.ToggleParticipantMicRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) EndConference(ctx context.Context, request models.EndConferenceRequest) (*pb.EndConferenceResponse, error) {
	res, err := c.Server.EndConference(ctx, &pb.EndConferenceRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
