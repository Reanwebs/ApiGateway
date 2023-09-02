package models

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

// type AcceptJoiningRequest struct {
// 	UserID       string `json:"userId"`
// 	ConferenceID string `json:"conferenceID"`
// }

// type DeclineJoiningRequest struct {
// 	UserID       string `json:"userId"`
// 	ConferenceID string `json:"conferenceID"`
// }

// type RemoveParticipantRequest struct {
// 	UserID       string `json:"userId"`
// 	ConferenceID string `json:"conferenceID"`
// 	Block        bool   `json:"block"`
// }

// type ToggleCameraRequest struct {
// 	UserID       string `json:"userID"`
// 	ConferenceID string `json:"conferenceID"`
// }

// type ToggleMicRequest struct {
// 	UserID       string `json:"userId"`
// 	ConferenceID string `json:"conferenceID"`
// }

// type ToggleParticipantCameraRequest struct {
// 	UserID       string `json:"userID"`
// 	ConferenceID string `json:"conferenceID"`
// }

//	type ToggleParticipantMicRequest struct {
//		UserID       string `json:"userID"`
//		ConferenceID string `json:"conferenceID"`
//	}
type HealthCheck struct {
	Data string `json:"data"`
}

type LeaveConferenceRequest struct {
	ConferenceID string `json:"conferenceID"`
}

type EndConferenceRequest struct {
	ConferenceID string `json:"conferenceID"`
}
type ScheduleConferenceRequest struct {
	Type             string `json:"type"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	Interest         string `json:"interest"`
	Recording        bool   `json:"recording"`
	Chat             bool   `json:"chat"`
	Broadcast        bool   `json:"broadcast"`
	Participantlimit int32  `json:"participantlimit"`
	Date             string `json:"date"`
	Time_seconds     int64  `json:"time_seconds"`
	Time_nanos       int32  `json:"time_nanos"`
}

// Start Conference

type StartPrivateConferenceRequest struct {
	Title            string `json:"title"`
	Description      string `json:"description"`
	Interest         string `json:"interest"`
	Recording        bool   `json:"recording"`
	Chat             bool   `json:"chat"`
	Broadcast        bool   `json:"broadcast"`
	Participantlimit int32  `json:"participantlimit"`
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
