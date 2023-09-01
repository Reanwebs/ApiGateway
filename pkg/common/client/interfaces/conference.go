package interfaces

import (
	"context"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb/conference"
)

type ConferenceClient interface {
	HealthCheck(context.Context, string, models.RetryConfig) (*conference.Response, error)
	ScheduleConference(context.Context, models.ScheduleConferenceRequest, models.RetryConfig) (*conference.ScheduleConferenceResponse, error)
	StartPrivateConference(context.Context, models.StartPrivateConferenceRequest, models.RetryConfig) (*conference.StartPrivateConferenceResponse, error)
	StartGroupConference(context.Context, models.StartGroupConferenceRequest, models.RetryConfig) (*conference.StartGroupConferenceResponse, error)
	StartPublicConference(context.Context, models.StartPublicConferenceRequest, models.RetryConfig) (*conference.StartPublicConferenceResponse, error)
	JoinPrivateConference(context.Context, models.JoinPrivateConferenceRequest, models.RetryConfig) (*conference.JoinPrivateConferenceResponse, error)
	JoinGroupConfernce(context.Context, models.JoinGroupConferenceRequest, models.RetryConfig) (*conference.JoinGroupConferenceResponse, error)
	JoinPublicConference(context.Context, models.JoinPublicConferenceRequest, models.RetryConfig) (*conference.JoinPublicConferenceResponse, error)
	AcceptJoining(context.Context, models.AcceptJoiningRequest, models.RetryConfig) (*conference.AcceptJoiningResponse, error)
	DeclineJoining(context.Context, models.DeclineJoiningRequest, models.RetryConfig) (*conference.DeclineJoiningResponse, error)
	LeavePrivateConference(context.Context, models.LeavePrivateConferenceRequest, models.RetryConfig) (*conference.LeavePrivateConferenceResponse, error)
	LeaveGroupConference(context.Context, models.LeaveGroupConferenceRequest, models.RetryConfig) (*conference.LeaveGroupConferenceResponse, error)
	LeavePublicConference(context.Context, models.LeavePublicConferenceRequest, models.RetryConfig) (*conference.LeavePublicConferenceResponse, error)
	RemovePrivateParticipant(context.Context, models.RemovePrivateParticipantRequest, models.RetryConfig) (*conference.RemovePrivateParticipantResponse, error)
	RemoveGroupParticipant(context.Context, models.RemoveGroupParticipantRequest, models.RetryConfig) (*conference.RemoveGroupParticipantResponse, error)
	RemovePublicParticipant(context.Context, models.RemovePublicParticipantRequest, models.RetryConfig) (*conference.RemovePublicParticipantResponse, error)
	EndPrivateConference(context.Context, models.EndPrivateConferenceRequest, models.RetryConfig) (*conference.EndPrivateConferenceResponse, error)
	EndGroupConference(context.Context, models.EndGroupConferenceRequest, models.RetryConfig) (*conference.EndGroupConferenceResponse, error)
	EndPublicConference(context.Context, models.EndPublicConferenceRequest, models.RetryConfig) (*conference.EndPublicConferenceResponse, error)
	ToggleCamera(context.Context, models.ToggleCameraRequest) (*conference.ToggleCameraResponse, error)
	ToggleMic(context.Context, models.ToggleMicRequest) (*conference.ToggleMicResponse, error)
	ToggleParticipantCamera(context.Context, models.ToggleParticipantCameraRequest) (*conference.ToggleParticipantCameraResponse, error)
	ToggleParticipantMic(context.Context, models.ToggleParticipantMicRequest) (*conference.ToggleParticipantMicResponse, error)
}
