package models

import "time"

type StartConferenceRequest struct {
	Type             string `json:"type"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	Interest         string `json:"interest"`
	Recording        bool   `json:"recording"`
	Chat             bool   `json:"chat"`
	Broadcast        bool   `json:"broadcast"`
	Participantlimit int32  `json:"participantlimit"`
}

type JoinConferenceRequest struct {
	ConferenceID string `json:"conferenceID"`
}

type HealthCheck struct {
	Data string `json:"data"`
}

type LeaveConferenceRequest struct {
	ConferenceID string `json:"conferenceID"`
}

type EndConferenceRequest struct {
	ConferenceID string `json:"conferenceID"`
}

// Start Conference

type StartPrivateConferenceRequest struct {
	Title            string `json:"title"`
	Description      string `json:"description"`
	Interest         string `json:"interest"`
	Recording        bool   `json:"recording"`
	Chat             bool   `json:"chat"`
	Broadcast        bool   `json:"broadcast"`
	Participantlimit string `json:"participantlimit"`
	SdpOffer         string `json:"sdpOffer"`
	ScheduledID      string `json:"scheduledID"`
}

type StartGroupConferenceRequest struct {
	GroupID          string `json:"groupID"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	Interest         string `json:"interest"`
	Recording        bool   `json:"recording"`
	Chat             bool   `json:"chat"`
	Broadcast        bool   `json:"broadcast"`
	Participantlimit int32  `json:"participantlimit"`
	SdpOffer         string `json:"sdpOffer"`
}

type StartPublicConferenceRequest struct {
	Title            string `json:"title"`
	Description      string `json:"description"`
	Interest         string `json:"interest"`
	JoinType         string `json:"joinType"`
	Recording        bool   `json:"recording"`
	Chat             bool   `json:"chat"`
	Broadcast        bool   `json:"broadcast"`
	Participantlimit int32  `json:"participantlimit"`
	SdpOffer         string `json:"sdpOffer"`
}

//  Join Conference

type JoinPrivateConferenceRequest struct {
	ConferenceID string `json:"conferenceID"`
}

type JoinGroupConferenceRequest struct {
	GroupID      string `json:"groupID"`
	ConferenceID string `json:"conferenceID"`
}

type JoinPublicConferenceRequest struct {
	ConferenceID string `json:"conferenceID"`
}

//  AcceptJoinRequest

type AcceptJoiningRequest struct {
	HostID       string `json:"hostID"`
	ConferenceID string `json:"conferenceID"`
}

// Decline Joining
type DeclineJoiningRequest struct {
	ConferenceID string `json:"conferenceID"`
}

// RemoveParticipant

type RemovePrivateParticipantRequest struct {
	HostID       string `json:"hostID"`
	ConferenceID string `json:"conferenceID"`
	Block        bool   `json:"block"`
}

type RemoveGroupParticipantRequest struct {
	HostID       string `json:"hostID"`
	ConferenceID string `json:"conferenceID"`
	Block        bool   `json:"block"`
}

type RemovePublicParticipantRequest struct {
	HostID       string `json:"hostID"`
	ConferenceID string `json:"conferenceID"`
	Block        bool   `json:"block"`
}

// ToggleCamera

type ToggleCameraRequest struct {
	ConferenceID string `json:"conferenceID"`
}

// ToggleMic

type ToggleMicRequest struct {
	ConferenceID string `json:"conferenceID"`
}

// ToggleParticipantCamera

type ToggleParticipantCameraRequest struct {
	ConferenceID string `json:"conferenceID"`
}

// ToggleParticipantMic

type ToggleParticipantMicRequest struct {
	ConferenceID string `json:"conferenceID"`
}

// LeaveConference

type LeavePrivateConferenceRequest struct {
	ConferenceID string `json:"conferenceID"`
}

type LeaveGroupConferenceRequest struct {
	ConferenceID string `json:"conferenceID"`
}

type LeavePublicConferenceRequest struct {
	ConferenceID string `json:"conferenceID"`
}

// EndConference

type EndPrivateConferenceRequest struct {
	ConferenceID string `json:"conferenceID"`
}

type EndGroupConferenceRequest struct {
	ConferenceID string `json:"conferenceID"`
}

type EndPublicConferenceRequest struct {
	ConferenceID string `json:"conferenceID"`
}
type SchedulePrivateConferenceRequest struct {
	Title            string `json:"title"`
	Description      string `json:"description"`
	Interest         string `json:"interest"`
	Recording        bool   `json:"recording"`
	Chat             bool   `json:"chat"`
	Participantlimit string `json:"participantlimit"`
	Time             string `json:"time"`
	Status           string `json:"status"`
	Duration         string `json:"duration"`
}

type ScheduleGroupConferenceRequest struct {
	GroupID          string `json:"groupID"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	Interest         string `json:"interest"`
	Recording        bool   `json:"recording"`
	Chat             bool   `json:"chat"`
	Participantlimit string `json:"Participantlimit"`
	Time             string `json:"time"`
	Status           string `json:"status"`
	Duration         string `json:"duration"`
}

type SchedulePublicConferenceRequest struct {
	Title            string `json:"title"`
	Description      string `json:"description"`
	Interest         string `json:"interest"`
	Recording        bool   `json:"recording"`
	Chat             bool   `json:"chat"`
	Participantlimit string `json:"participantlimit"`
	Time             string `json:"time"`
	Status           string `json:"status"`
	Duration         string `json:"duration"`
}
type ScheduledConference struct {
	UserID           string
	ScheduleID       string
	Title            string
	Description      string
	Interest         string
	Chat             bool
	Participantlimit int32
	Time             time.Time
	Status           string
	Durations        int32
}

type ScheduledConferenceResponse struct {
	Result              string
	ScheduledConference []ScheduledConference
	Err                 error
}

type StartStreamRequest struct {
	Title       string `json:"title"`
	Discription string `json:"discription"`
	Interest    string `json:"interest"`
	ThubnailID  string `json:"thubnailID"`
}

type JoinStreamRequest struct {
	StreamID string `json:"streamID"`
}

type LeaveStreamRequest struct {
	StreamID string `json:"streamID"`
}

type StopStreamRequest struct {
	StreamID string `json:"streamID"`
}
