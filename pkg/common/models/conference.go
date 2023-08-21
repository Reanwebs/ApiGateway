package models

type StartConferenceRequest struct {
	UserID           string `json:"userID"`
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
	UserID       string `json:"userId"`
	ConferenceID int32  `json:"conferenceID"`
}

type AcceptJoiningRequest struct {
	UserID       string `json:"userId"`
	ConferenceID int32  `json:"conferenceID"`
}

type DeclineJoiningRequest struct {
	UserID       string `json:"userId"`
	ConferenceID int32  `json:"conferenceID"`
}

type RemoveParticipantRequest struct {
	UserID       string `json:"userId"`
	ConferenceID int32  `json:"conferenceID"`
	Block        bool   `json:"block"`
}

type ToggleCameraRequest struct {
	UserID       string `json:"userID"`
	ConferenceID int32  `json:"conferenceID"`
}

type ToggleMicRequest struct {
	UserID       string `json:"userId"`
	ConferenceID int32  `json:"conferenceID"`
}

type ToggleParticipantCameraRequest struct {
	UserID       string `json:"userID"`
	ConferenceID int32  `json:"conferenceID"`
}

type ToggleParticipantMicRequest struct {
	UserID       string `json:"userID"`
	ConferenceID int32  `json:"conferenceID"`
}

type LeaveConferenceRequest struct {
	UserID       string `json:"userID"`
	ConferenceID int32  `json:"conferenceID"`
}

type EndConferenceRequest struct {
	UserID       string `json:"userID"`
	ConferenceID int32  `json:"conferenceID"`
}
