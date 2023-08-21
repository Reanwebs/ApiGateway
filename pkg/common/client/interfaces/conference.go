package interfaces

import (
	"context"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb"
)

type ConferenceClient interface {
	HealthCheck(context.Context, string) (*pb.Response, error)
	StartConference(context.Context, models.StartConferenceRequest) (*pb.StartConferenceResponse, error)
	JoinConference(context.Context, models.JoinConferenceRequest) (*pb.JoinConferenceResponse, error)
	AcceptJoining(context.Context, models.AcceptJoiningRequest) (*pb.AcceptJoiningResponse, error)
	DeclineJoining(context.Context, models.DeclineJoiningRequest) (*pb.DeclineJoiningResponse, error)
	LeaveConference(context.Context, models.LeaveConferenceRequest) (*pb.LeaveConferenceResponse, error)
	RemoveParticipant(context.Context, models.RemoveParticipantRequest) (*pb.RemoveParticipantResponse, error)
	ToggleCamera(context.Context, models.ToggleCameraRequest) (*pb.ToggleCameraResponse, error)
	ToggleMic(context.Context, models.ToggleMicRequest) (*pb.ToggleMicResponse, error)
	ToggleParticipantCamera(context.Context, models.ToggleParticipantCameraRequest) (*pb.ToggleParticipantCameraResponse, error)
	ToggleParticipantMic(context.Context, models.ToggleParticipantMicRequest) (*pb.ToggleParticipantMicResponse, error)
	EndConference(context.Context, models.EndConferenceRequest) (*pb.EndConferenceResponse, error)
}
