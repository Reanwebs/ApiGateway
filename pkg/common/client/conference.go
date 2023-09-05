package client

import (
	"context"
	"errors"
	"fmt"
	"gateway/pkg/common/client/interfaces"
	"gateway/pkg/common/config"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb/conference"
	"gateway/pkg/utils"
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

func (c *conferenceClient) HealthCheck(ctx context.Context, request string, retryConfig models.RetryConfig) (*conference.Response, error) {

	var res *conference.Response
	var err error
	startTime := time.Now()

	for retryCount := 1; retryCount <= retryConfig.MaxRetries; retryCount++ {
		fmt.Println("try \t=>\t", retryCount)
		if time.Since(startTime) > retryConfig.MaxDuration {
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

			// Check if the error is non-retryable
			if utils.IsNonRetryableError(err) {
				return nil, err
			}
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryConfig.RetryInterval):
		}
	}

	return nil, err
}

func (c *conferenceClient) ScheduleConference(ctx context.Context, request models.ScheduleConferenceRequest, retryConfig models.RetryConfig) (*conference.ScheduleConferenceResponse, error) {
	var res *conference.ScheduleConferenceResponse
	var err error

	// Retrieve the "userId" from the context.
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	startTime := time.Now()

	for retryCount := 1; retryCount <= retryConfig.MaxRetries; retryCount++ {
		fmt.Println("try.........", retryCount)

		if time.Since(startTime) > retryConfig.MaxDuration {
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
			res, err = c.Server.ScheduleConference(ctx1, &conference.ScheduleConferenceRequest{
				UserID:           userId,
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
			cancel()

			if err == nil {
				return res, nil
			}

			// Check if the error is non-retryable
			if utils.IsNonRetryableError(err) {
				return nil, err
			}
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryConfig.RetryInterval):
		}
	}

	return nil, err
}

func (c *conferenceClient) StartPrivateConference(ctx context.Context, request models.StartPrivateConferenceRequest, retryConfig models.RetryConfig) (*conference.StartPrivateConferenceResponse, error) {
	var res *conference.StartPrivateConferenceResponse
	var err error

	// Retrieve the "userId" from the context.
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	startTime := time.Now()

	for retryCount := 1; retryCount <= retryConfig.MaxRetries; retryCount++ {
		fmt.Println("try.........", retryCount)

		if time.Since(startTime) > retryConfig.MaxDuration {
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
			res, err = c.Server.StartPrivateConference(ctx1, &conference.StartPrivateConferenceRequest{
				UserID:           userId,
				Title:            request.Title,
				Description:      request.Description,
				Interest:         request.Interest,
				Recording:        request.Recording,
				Chat:             request.Chat,
				Broadcast:        request.Broadcast,
				Participantlimit: request.Participantlimit,
			})
			cancel()

			if err == nil {
				return res, nil
			}

			// Check if the error is non-retryable
			if utils.IsNonRetryableError(err) {
				return nil, err
			}
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryConfig.RetryInterval):
		}
	}

	return nil, err

}

func (c *conferenceClient) StartGroupConference(ctx context.Context, request models.StartGroupConferenceRequest, retryConfig models.RetryConfig) (*conference.StartGroupConferenceResponse, error) {
	var res *conference.StartGroupConferenceResponse
	var err error

	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	startTime := time.Now()

	for retryCount := 1; retryCount <= retryConfig.MaxRetries; retryCount++ {
		fmt.Println("try.........", retryCount)

		if time.Since(startTime) > retryConfig.MaxDuration {
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
			res, err = c.Server.StartGroupConference(ctx1, &conference.StartGroupConferenceRequest{
				UserID:           userId,
				GroupID:          request.GroupID,
				Title:            request.Title,
				Description:      request.Description,
				Interest:         request.Interest,
				Recording:        request.Recording,
				Chat:             request.Chat,
				Broadcast:        request.Broadcast,
				Participantlimit: request.Participantlimit,
			})
			cancel()

			if err == nil {
				return res, nil
			}

			// Check if the error is non-retryable
			if utils.IsNonRetryableError(err) {
				return nil, err
			}

		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryConfig.RetryInterval):
		}

	}
	return nil, err
}

func (c *conferenceClient) StartPublicConference(ctx context.Context, request models.StartPublicConferenceRequest, retryConfig models.RetryConfig) (*conference.StartPublicConferenceResponse, error) {
	var res *conference.StartPublicConferenceResponse
	var err error

	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	startTime := time.Now()

	for retryCount := 1; retryCount <= retryConfig.MaxRetries; retryCount++ {
		fmt.Println("try.........", retryCount)

		if time.Since(startTime) > retryConfig.MaxDuration {
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
			res, err = c.Server.StartPublicConference(ctx1, &conference.StartPublicConferenceRequest{
				UserID:           userId,
				Title:            request.Title,
				Description:      request.Description,
				Interest:         request.Interest,
				JoinType:         request.JoinType,
				Recording:        request.Recording,
				Chat:             request.Chat,
				Broadcast:        request.Broadcast,
				Participantlimit: request.Participantlimit,
			})
			cancel()

			if err == nil {
				return res, nil
			}

			// Check if the error is non-retryable
			if utils.IsNonRetryableError(err) {
				return nil, err
			}
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryConfig.RetryInterval):
		}

	}
	return nil, err

}

func (c *conferenceClient) JoinPrivateConference(ctx context.Context, request models.JoinPrivateConferenceRequest, retryConfig models.RetryConfig) (*conference.JoinPrivateConferenceResponse, error) {
	var res *conference.JoinPrivateConferenceResponse
	var err error

	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	startTime := time.Now()

	for retryCount := 1; retryCount <= retryConfig.MaxRetries; retryCount++ {
		fmt.Println("try.........", retryCount)

		if time.Since(startTime) > retryConfig.MaxDuration {
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
			res, err = c.Server.JoinPrivateConference(ctx1, &conference.JoinPrivateConferenceRequest{
				UserID:       userId,
				ConferenceID: request.ConferenceID,
			})
			cancel()

			if err == nil {
				return res, nil
			}

			// Check if the error is non-retryable
			if utils.IsNonRetryableError(err) {
				return nil, err
			}
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryConfig.RetryInterval):
		}
	}
	return nil, err
}

func (c *conferenceClient) JoinGroupConfernce(ctx context.Context, request models.JoinGroupConferenceRequest, retryConfig models.RetryConfig) (*conference.JoinGroupConferenceResponse, error) {
	var res *conference.JoinGroupConferenceResponse
	var err error

	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	startTime := time.Now()

	for retryCount := 1; retryCount <= retryConfig.MaxRetries; retryCount++ {
		fmt.Println("try.........", retryCount)

		if time.Since(startTime) > retryConfig.MaxDuration {
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
			res, err = c.Server.JoinGroupConfernce(ctx1, &conference.JoinGroupConferenceRequest{
				UserID:       userId,
				GroupID:      request.GroupID,
				ConferenceID: request.ConferenceID,
			})
			cancel()

			if err == nil {
				return res, nil
			}

			// Check if the error is non-retryable
			if utils.IsNonRetryableError(err) {
				return nil, err
			}
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryConfig.RetryInterval):
		}
	}

	return nil, err
}

func (c *conferenceClient) JoinPublicConference(ctx context.Context, request models.JoinPublicConferenceRequest, retryConfig models.RetryConfig) (*conference.JoinPublicConferenceResponse, error) {
	var res *conference.JoinPublicConferenceResponse
	var err error

	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	startTime := time.Now()

	for retryCount := 1; retryCount <= retryConfig.MaxRetries; retryCount++ {
		fmt.Println("try.........", retryCount)

		if time.Since(startTime) > retryConfig.MaxDuration {
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
			res, err = c.Server.JoinPublicConference(ctx1, &conference.JoinPublicConferenceRequest{
				UserID:       userId,
				ConferenceID: request.ConferenceID,
			})
			cancel()

			if err == nil {
				return res, nil
			}

			// Check if the error is non-retryable
			if utils.IsNonRetryableError(err) {
				return nil, err
			}
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryConfig.RetryInterval):
		}
	}

	return nil, err
}

func (c *conferenceClient) AcceptJoining(ctx context.Context, request models.AcceptJoiningRequest, retryConfig models.RetryConfig) (*conference.AcceptJoiningResponse, error) {
	var res *conference.AcceptJoiningResponse
	var err error

	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	startTime := time.Now()

	for retryCount := 1; retryCount <= retryConfig.MaxRetries; retryCount++ {
		fmt.Println("try.........", retryCount)

		if time.Since(startTime) > retryConfig.MaxDuration {
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
			res, err = c.Server.AcceptJoining(ctx1, &conference.AcceptJoiningRequest{
				UserID:       userId,
				ConferenceID: request.ConferenceID,
			})
			cancel()

			if err == nil {
				return res, nil
			}

			// Check if the error is non-retryable
			if utils.IsNonRetryableError(err) {
				return nil, err
			}
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryConfig.RetryInterval):
		}

	}

	return nil, err
}

func (c *conferenceClient) DeclineJoining(ctx context.Context, request models.DeclineJoiningRequest, retryConfig models.RetryConfig) (*conference.DeclineJoiningResponse, error) {
	var res *conference.DeclineJoiningResponse
	var err error

	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	startTime := time.Now()

	for retryCount := 1; retryCount <= retryConfig.MaxRetries; retryCount++ {
		fmt.Println("try.........", retryCount)

		if time.Since(startTime) > retryConfig.MaxDuration {
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
			res, err = c.Server.DeclineJoining(ctx1, &conference.DeclineJoiningRequest{
				UserID:       userId,
				ConferenceID: request.ConferenceID,
			})
			cancel()

			if err == nil {
				return res, nil
			}

			// Check if the error is non-retryable
			if utils.IsNonRetryableError(err) {
				return nil, err
			}
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryConfig.RetryInterval):
		}
	}

	return nil, err
}

func (c *conferenceClient) LeavePrivateConference(ctx context.Context, request models.LeavePrivateConferenceRequest, retryConfig models.RetryConfig) (*conference.LeavePrivateConferenceResponse, error) {
	var res *conference.LeavePrivateConferenceResponse
	var err error

	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	startTime := time.Now()

	for retryCount := 1; retryCount <= retryConfig.MaxRetries; retryCount++ {
		fmt.Println("try.........", retryCount)

		if time.Since(startTime) > retryConfig.MaxDuration {
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
			res, err = c.Server.LeavePrivateConference(ctx1, &conference.LeavePrivateConferenceRequest{
				UserID:       userId,
				ConferenceID: request.ConferenceID,
			})
			cancel()

			if err == nil {
				return res, nil
			}

			// Check if the error is non-retryable
			if utils.IsNonRetryableError(err) {
				return nil, err
			}
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryConfig.RetryInterval):
		}
	}

	return nil, err
}

func (c *conferenceClient) LeaveGroupConference(ctx context.Context, request models.LeaveGroupConferenceRequest, retryConfig models.RetryConfig) (*conference.LeaveGroupConferenceResponse, error) {
	var res *conference.LeaveGroupConferenceResponse
	var err error

	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	startTime := time.Now()

	for retryCount := 1; retryCount <= retryConfig.MaxRetries; retryCount++ {
		fmt.Println("try.........", retryCount)

		if time.Since(startTime) > retryConfig.MaxDuration {
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
			res, err = c.Server.LeaveGroupConference(ctx1, &conference.LeaveGroupConferenceRequest{
				UserID:       userId,
				ConferenceID: request.ConferenceID,
			})
			cancel()

			if err == nil {
				return res, nil
			}

			// Check if the error is non-retryable
			if utils.IsNonRetryableError(err) {
				return nil, err
			}
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryConfig.RetryInterval):
		}
	}

	return nil, err
}

func (c *conferenceClient) LeavePublicConference(ctx context.Context, request models.LeavePublicConferenceRequest, retryConfig models.RetryConfig) (*conference.LeavePublicConferenceResponse, error) {
	var res *conference.LeavePublicConferenceResponse
	var err error

	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	startTime := time.Now()

	for retryCount := 1; retryCount <= retryConfig.MaxRetries; retryCount++ {
		fmt.Println("try.........", retryCount)

		if time.Since(startTime) > retryConfig.MaxDuration {
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
			res, err = c.Server.LeavePublicConference(ctx1, &conference.LeavePublicConferenceRequest{
				UserID:       userId,
				ConferenceID: request.ConferenceID,
			})
			cancel()

			if err == nil {
				return res, nil
			}

			// Check if the error is non-retryable
			if utils.IsNonRetryableError(err) {
				return nil, err
			}
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryConfig.RetryInterval):
		}
	}

	return nil, err
}

func (c *conferenceClient) RemovePrivateParticipant(ctx context.Context, request models.RemovePrivateParticipantRequest, retryConfig models.RetryConfig) (*conference.RemovePrivateParticipantResponse, error) {
	var res *conference.RemovePrivateParticipantResponse
	var err error

	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	startTime := time.Now()

	for retryCount := 1; retryCount <= retryConfig.MaxRetries; retryCount++ {
		fmt.Println("try.........", retryCount)

		if time.Since(startTime) > retryConfig.MaxDuration {
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
			res, err = c.Server.RemovePrivateParticipant(ctx1, &conference.RemovePrivateParticipantRequest{
				UserID:       userId,
				HostID:       request.HostID,
				ConferenceID: request.ConferenceID,
				Block:        request.Block,
			})
			cancel()

			if err == nil {
				return res, nil
			}

			// Check if the error is non-retryable
			if utils.IsNonRetryableError(err) {
				return nil, err
			}
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryConfig.RetryInterval):
		}
	}

	return nil, err
}

func (c *conferenceClient) RemoveGroupParticipant(ctx context.Context, request models.RemoveGroupParticipantRequest, retryConfig models.RetryConfig) (*conference.RemoveGroupParticipantResponse, error) {
	var res *conference.RemoveGroupParticipantResponse
	var err error

	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	startTime := time.Now()

	for retryCount := 1; retryCount <= retryConfig.MaxRetries; retryCount++ {
		fmt.Println("try.........", retryCount)

		if time.Since(startTime) > retryConfig.MaxDuration {
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
			res, err = c.Server.RemoveGroupParticipant(ctx1, &conference.RemoveGroupParticipantRequest{
				UserID:       userId,
				HostID:       request.HostID,
				ConferenceID: request.ConferenceID,
				Block:        request.Block,
			})
			cancel()

			if err == nil {
				return res, nil
			}

			// Check if the error is non-retryable
			if utils.IsNonRetryableError(err) {
				return nil, err
			}
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryConfig.RetryInterval):
		}
	}

	return res, nil
}

func (c *conferenceClient) RemovePublicParticipant(ctx context.Context, request models.RemovePublicParticipantRequest, retryConfig models.RetryConfig) (*conference.RemovePublicParticipantResponse, error) {
	var res *conference.RemovePublicParticipantResponse
	var err error

	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	startTime := time.Now()

	for retryCount := 1; retryCount <= retryConfig.MaxRetries; retryCount++ {
		fmt.Println("try.........", retryCount)

		if time.Since(startTime) > retryConfig.MaxDuration {
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
			res, err = c.Server.RemovePublicParticipant(ctx1, &conference.RemovePublicParticipantRequest{
				UserID:       userId,
				HostID:       request.HostID,
				ConferenceID: request.ConferenceID,
				Block:        request.Block,
			})
			cancel()

			if err == nil {
				return res, nil
			}

			// Check if the error is non-retryable
			if utils.IsNonRetryableError(err) {
				return nil, err
			}
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryConfig.RetryInterval):
		}
	}

	return nil, err
}

func (c *conferenceClient) EndPrivateConference(ctx context.Context, request models.EndPrivateConferenceRequest, retryConfig models.RetryConfig) (*conference.EndPrivateConferenceResponse, error) {
	var res *conference.EndPrivateConferenceResponse
	var err error

	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	startTime := time.Now()

	for retryCount := 1; retryCount <= retryConfig.MaxRetries; retryCount++ {
		fmt.Println("try.........", retryCount)

		if time.Since(startTime) > retryConfig.MaxDuration {
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
			res, err = c.Server.EndPrivateConference(ctx1, &conference.EndPrivateConferenceRequest{
				UserID:       userId,
				ConferenceID: request.ConferenceID,
			})
			cancel()

			if err == nil {
				return res, nil
			}

			// Check if the error is non-retryable
			if utils.IsNonRetryableError(err) {
				return nil, err
			}
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryConfig.RetryInterval):
		}
	}

	return res, nil
}

func (c *conferenceClient) EndGroupConference(ctx context.Context, request models.EndGroupConferenceRequest, retryConfig models.RetryConfig) (*conference.EndGroupConferenceResponse, error) {
	var res *conference.EndGroupConferenceResponse
	var err error

	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	startTime := time.Now()

	for retryCount := 1; retryCount <= retryConfig.MaxRetries; retryCount++ {
		fmt.Println("try.........", retryCount)

		if time.Since(startTime) > retryConfig.MaxDuration {
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
			res, err = c.Server.EndGroupConference(ctx1, &conference.EndGroupConferenceRequest{
				UserID:       userId,
				ConferenceID: request.ConferenceID,
			})
			cancel()

			if err == nil {
				return res, nil
			}

			// Check if the error is non-retryable
			if utils.IsNonRetryableError(err) {
				return nil, err
			}
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryConfig.RetryInterval):
		}
	}

	return nil, err
}

func (c *conferenceClient) EndPublicConference(ctx context.Context, request models.EndPublicConferenceRequest, retryConfig models.RetryConfig) (*conference.EndPublicConferenceResponse, error) {
	var res *conference.EndPublicConferenceResponse
	var err error

	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	startTime := time.Now()

	for retryCount := 1; retryCount <= retryConfig.MaxRetries; retryCount++ {
		fmt.Println("try.........", retryCount)

		if time.Since(startTime) > retryConfig.MaxDuration {
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
			res, err = c.Server.EndPublicConference(ctx1, &conference.EndPublicConferenceRequest{
				UserID:       userId,
				ConferenceID: request.ConferenceID,
			})
			cancel()

			if err == nil {
				return res, nil
			}

			// Check if the error is non-retryable
			if utils.IsNonRetryableError(err) {
				return nil, err
			}
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryConfig.RetryInterval):
		}
	}

	return nil, err
}

func (c *conferenceClient) ToggleCamera(ctx context.Context, request models.ToggleCameraRequest) (*conference.ToggleCameraResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := c.Server.ToggleCamera(ctx, &conference.ToggleCameraRequest{
		UserID:       userId,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) ToggleMic(ctx context.Context, request models.ToggleMicRequest) (*conference.ToggleMicResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := c.Server.ToggleMic(ctx, &conference.ToggleMicRequest{
		UserID:       userId,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) ToggleParticipantCamera(ctx context.Context, request models.ToggleParticipantCameraRequest) (*conference.ToggleParticipantCameraResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := c.Server.ToggleParticipantCamera(ctx, &conference.ToggleParticipantCameraRequest{
		UserID:       userId,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *conferenceClient) ToggleParticipantMic(ctx context.Context, request models.ToggleParticipantMicRequest) (*conference.ToggleParticipantMicResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	res, err := c.Server.ToggleParticipantMic(ctx, &conference.ToggleParticipantMicRequest{
		UserID:       userId,
		ConferenceID: request.ConferenceID,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
