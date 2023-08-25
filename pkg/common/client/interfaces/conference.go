package interfaces

import (
	"context"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb/conference"
)

type ConferenceClient interface {
	HealthCheck(context.Context, string) (*conference.Response, error)
	ScheduleConference(context.Context, models.ScheduleConferenceRequest) (*conference.ScheduleConferenceResponse, error)
	StartPrivateConference(context.Context, models.StartPrivateConferenceRequest) (*conference.StartPrivateConferenceResponse, error)
	StartGroupConference(context.Context, models.StartGroupConferenceRequest) (*conference.StartGroupConferenceResponse, error)
	StartPublicConference(context.Context, models.StartPublicConferenceRequest) (*conference.StartPublicConferenceResponse, error)
	JoinPrivateConference(context.Context, models.JoinPrivateConferenceRequest) (*conference.JoinPrivateConferenceResponse, error)
	JoinGroupConfernce(context.Context, models.JoinGroupConferenceRequest) (*conference.JoinGroupConferenceResponse, error)
	JoinPublicConference(context.Context, models.JoinPublicConferenceRequest) (*conference.JoinPublicConferenceResponse, error)
	AcceptJoining(context.Context, models.AcceptJoiningRequest) (*conference.AcceptJoiningResponse, error)
	DeclineJoining(context.Context, models.DeclineJoiningRequest) (*conference.DeclineJoiningResponse, error)
	LeavePrivateConference(context.Context, models.LeavePrivateConferenceRequest) (*conference.LeavePrivateConferenceResponse, error)
	LeaveGroupConference(context.Context, models.LeaveGroupConferenceRequest) (*conference.LeaveGroupConferenceResponse, error)
	LeavePublicConference(context.Context, models.LeavePublicConferenceRequest) (*conference.LeavePublicConferenceResponse, error)
	RemovePrivateParticipant(context.Context, models.RemovePrivateParticipantRequest) (*conference.RemovePrivateParticipantResponse, error)
	RemoveGroupParticipant(context.Context, models.RemoveGroupParticipantRequest) (*conference.RemoveGroupParticipantResponse, error)
	RemovePublicParticipant(context.Context, models.RemovePublicParticipantRequest) (*conference.RemovePublicParticipantResponse, error)
	EndPrivateConference(context.Context, models.EndPrivateConferenceRequest) (*conference.EndPrivateConferenceResponse, error)
	EndGroupConference(context.Context, models.EndGroupConferenceRequest) (*conference.EndGroupConferenceResponse, error)
	EndPublicConference(context.Context, models.EndPublicConferenceRequest) (*conference.EndPublicConferenceResponse, error)
	ToggleCamera(context.Context, models.ToggleCameraRequest) (*conference.ToggleCameraResponse, error)
	ToggleMic(context.Context, models.ToggleMicRequest) (*conference.ToggleMicResponse, error)
	ToggleParticipantCamera(context.Context, models.ToggleParticipantCameraRequest) (*conference.ToggleParticipantCameraResponse, error)
	ToggleParticipantMic(context.Context, models.ToggleParticipantMicRequest) (*conference.ToggleParticipantMicResponse, error)
}
