package models

import "github.com/golang/protobuf/ptypes/timestamp"

type Participants struct {
	ConferenceID     string
	ParticipantID    int
	IsOrganizer      bool
	MicPermission    bool
	CameraPermission bool
	JoinTime         timestamp.Timestamp
	LeaveTime        timestamp.Timestamp
}

type Conference struct {
	ConferenceID     string
	OrganizerID      string
	Type             string
	StartTime        timestamp.Timestamp
	EndTime          timestamp.Timestamp
	ChatPermission   bool
	Broadcast        bool
	ParticipantLimit uint
	Title            string
	Description      string
	Link             string
	Interest         string
}

type ConferencePublicChat struct {
	ConferenceID  string
	ParticipantID int
	Message       string
	Time          timestamp.Timestamp
}

type ConferencePrivateChat struct {
	ConferenceID  string
	ParticipantID int
	RecieverID    int
	Message       string
	Time          timestamp.Timestamp
}

type StartConference struct {
	UserID           string
	ConferenceType   string
	Title            string
	Description      string
	Interest         string
	Recording        bool
	Chat             bool
	Broadcast        bool
	ParticipantLimit uint
}
