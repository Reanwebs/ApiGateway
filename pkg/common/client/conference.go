package client

import (
	"context"
	"errors"
	"fmt"
	"gateway/pkg/common/client/interfaces"
	"gateway/pkg/common/config"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb/conference"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type conferenceClient struct {
	Server conference.ConferenceClient
}

func InitConferenceClient(c *config.Config) (interfaces.ConferenceClient, error) {
	cc, err := grpc.Dial(c.ConferenceService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return NewConferenceClient(conference.NewConferenceClient(cc)), nil
}

func NewConferenceClient(server conference.ConferenceClient) interfaces.ConferenceClient {
	return &conferenceClient{
		Server: server,
	}
}

func (c *conferenceClient) HealthCheck(ctx context.Context, request string) (*conference.Response, error) {
	var res *conference.Response
	var err error

	maxRetries := 5
	maxDuration := 5 * time.Second
	startTime := time.Now()

	for retryCount := 1; retryCount <= maxRetries; retryCount++ {
		if time.Since(startTime) > maxDuration {
			err = errors.New("time limit exceeded")
			break
		}

		ctx1, cancel := context.WithTimeout(ctx, 30*time.Second)

		done := make(chan struct{})
		go func() {
			defer close(done)
			<-ctx.Done()
		}()

		select {
		case <-done:
			cancel()
			return nil, context.Canceled

		default:
			res, err = c.Server.HealthCheck(ctx1, &conference.Request{
				Data: request,
			})
			cancel()
			if err == nil {
				return res, nil
			}
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(time.Second):
		}
	}

	return nil, err

}

// func (c *conferenceClient) ScheduleConference(ctx context.Context, request models.ScheduleConferenceRequest) (*conference.ScheduleConferenceResponse, error) {
// 	ctx1, _ := context.WithTimeout(ctx, 30*time.Second)
// 	go func() {
// 		<-ctx.Done()
// 		// log golang done channel patterns, golang generated channel pattern
// 	}()
// 	res, err := c.Server.ScheduleConference(ctx1, &conference.ScheduleConferenceRequest{
// 		UserID:           request.UserID,
// 		Type:             request.Type,
// 		Title:            request.Type,
// 		Description:      request.Description,
// 		Interest:         request.Interest,
// 		Recording:        request.Recording,
// 		Chat:             request.Chat,
// 		Broadcast:        request.Broadcast,
// 		Participantlimit: request.Participantlimit,
// 		Date:             request.Date,
// 		TimeSeconds:      request.Time_seconds,
// 		TimeNanos:        request.Time_nanos,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

func (c *conferenceClient) ScheduleConference(ctx context.Context, request models.ScheduleConferenceRequest) (*conference.ScheduleConferenceResponse, error) {
	var res *conference.ScheduleConferenceResponse
	var err error
	maxRetries := 3
	for i := 0; i <= maxRetries; i++ {
		fmt.Println("try.........", i)
		ctx1, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		done := make(chan struct{})
		go func() {
			defer close(done)
			<-ctx.Done() // Listen for cancellation
		}()

		select {
		case <-done:
			// Handle cancellation
			return nil, context.Canceled
		default:
			res, err = c.Server.ScheduleConference(ctx1, &conference.ScheduleConferenceRequest{
				UserID:           request.UserID,
				Type:             request.Type,
				Title:            request.Title,
				Description:      request.Description,
				Interest:         request.Interest,
				Recording:        request.Recording,
				Chat:             request.Chat,
				Broadcast:        request.Broadcast,
				Participantlimit: request.Participantlimit,
				Date:             request.Date,
				TimeSeconds:      3000,
				TimeNanos:        500,
			})
			if err == nil {
				return res, nil // Request succeeded, return response
			}
			// Request failed, continue to retry if possible
		}

		select {
		case <-ctx.Done():
			// Handle timeout
			return nil, context.DeadlineExceeded
		case <-time.After(time.Second): // Wait before retrying
		}
	}

	return nil, err // Return last error after maxRetries
}

func (c *conferenceClient) StartPrivateConference(ctx context.Context, request models.StartPrivateConferenceRequest) (*conference.StartPrivateConferenceResponse, error) {
	res, err := c.Server.StartPrivateConference(ctx, &conference.StartPrivateConferenceRequest{
		UserID:           request.UserID,
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

func (c *conferenceClient) StartGroupConference(ctx context.Context, request models.StartGroupConferenceRequest) (*conference.StartGroupConferenceResponse, error) {
	res, err := c.Server.StartGroupConference(ctx, &conference.StartGroupConferenceRequest{
		UserID:           request.UserID,
		GroupID:          request.GroupID,
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

func (c *conferenceClient) StartPublicConference(ctx context.Context, request models.StartPublicConferenceRequest) (*conference.StartPublicConferenceResponse, error) {
	res, err := c.Server.StartPublicConference(ctx, &conference.StartPublicConferenceRequest{
		UserID:           request.UserID,
		Title:            request.Title,
		Description:      request.Description,
		Interest:         request.Interest,
		JoinType:         request.JoinType,
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

func (c *conferenceClient) JoinPrivateConference(ctx context.Context, request models.JoinPrivateConferenceRequest) (*conference.JoinPrivateConferenceResponse, error) {
	res, err := c.Server.JoinPrivateConference(ctx, &conference.JoinPrivateConferenceRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) JoinGroupConfernce(ctx context.Context, request models.JoinGroupConferenceRequest) (*conference.JoinGroupConferenceResponse, error) {
	res, err := c.Server.JoinGroupConfernce(ctx, &conference.JoinGroupConferenceRequest{
		UserID:       request.UserID,
		GroupID:      request.GroupID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) JoinPublicConference(ctx context.Context, request models.JoinPublicConferenceRequest) (*conference.JoinPublicConferenceResponse, error) {
	res, err := c.Server.JoinPublicConference(ctx, &conference.JoinPublicConferenceRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) AcceptJoining(ctx context.Context, request models.AcceptJoiningRequest) (*conference.AcceptJoiningResponse, error) {
	res, err := c.Server.AcceptJoining(ctx, &conference.AcceptJoiningRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) DeclineJoining(ctx context.Context, request models.DeclineJoiningRequest) (*conference.DeclineJoiningResponse, error) {
	res, err := c.Server.DeclineJoining(ctx, &conference.DeclineJoiningRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) LeavePrivateConference(ctx context.Context, request models.LeavePrivateConferenceRequest) (*conference.LeavePrivateConferenceResponse, error) {
	res, err := c.Server.LeavePrivateConference(ctx, &conference.LeavePrivateConferenceRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) LeaveGroupConference(ctx context.Context, request models.LeaveGroupConferenceRequest) (*conference.LeaveGroupConferenceResponse, error) {
	res, err := c.Server.LeaveGroupConference(ctx, &conference.LeaveGroupConferenceRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) LeavePublicConference(ctx context.Context, request models.LeavePublicConferenceRequest) (*conference.LeavePublicConferenceResponse, error) {
	res, err := c.Server.LeavePublicConference(ctx, &conference.LeavePublicConferenceRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) RemovePrivateParticipant(ctx context.Context, request models.RemovePrivateParticipantRequest) (*conference.RemovePrivateParticipantResponse, error) {
	res, err := c.Server.RemovePrivateParticipant(ctx, &conference.RemovePrivateParticipantRequest{
		UserID:       request.UserID,
		HostID:       request.HostID,
		ConferenceID: request.ConferenceID,
		Block:        request.Block,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) RemoveGroupParticipant(ctx context.Context, request models.RemoveGroupParticipantRequest) (*conference.RemoveGroupParticipantResponse, error) {
	res, err := c.Server.RemoveGroupParticipant(ctx, &conference.RemoveGroupParticipantRequest{
		UserID:       request.UserID,
		HostID:       request.HostID,
		ConferenceID: request.ConferenceID,
		Block:        request.Block,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) RemovePublicParticipant(ctx context.Context, request models.RemovePublicParticipantRequest) (*conference.RemovePublicParticipantResponse, error) {
	res, err := c.Server.RemovePublicParticipant(ctx, &conference.RemovePublicParticipantRequest{
		UserID:       request.UserID,
		HostID:       request.HostID,
		ConferenceID: request.ConferenceID,
		Block:        request.Block,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) EndPrivateConference(ctx context.Context, request models.EndPrivateConferenceRequest) (*conference.EndPrivateConferenceResponse, error) {
	res, err := c.Server.EndPrivateConference(ctx, &conference.EndPrivateConferenceRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) EndGroupConference(ctx context.Context, request models.EndGroupConferenceRequest) (*conference.EndGroupConferenceResponse, error) {
	res, err := c.Server.EndGroupConference(ctx, &conference.EndGroupConferenceRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) EndPublicConference(ctx context.Context, request models.EndPublicConferenceRequest) (*conference.EndPublicConferenceResponse, error) {
	res, err := c.Server.EndPublicConference(ctx, &conference.EndPublicConferenceRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) ToggleCamera(ctx context.Context, request models.ToggleCameraRequest) (*conference.ToggleCameraResponse, error) {
	res, err := c.Server.ToggleCamera(ctx, &conference.ToggleCameraRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) ToggleMic(ctx context.Context, request models.ToggleMicRequest) (*conference.ToggleMicResponse, error) {
	res, err := c.Server.ToggleMic(ctx, &conference.ToggleMicRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) ToggleParticipantCamera(ctx context.Context, request models.ToggleParticipantCameraRequest) (*conference.ToggleParticipantCameraResponse, error) {
	res, err := c.Server.ToggleParticipantCamera(ctx, &conference.ToggleParticipantCameraRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) ToggleParticipantMic(ctx context.Context, request models.ToggleParticipantMicRequest) (*conference.ToggleParticipantMicResponse, error) {
	res, err := c.Server.ToggleParticipantMic(ctx, &conference.ToggleParticipantMicRequest{
		UserID:       request.UserID,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
