package models

type Request struct {
	Data string `json:"data"`
}

type UpdateWalletRequest struct {
	UserID   string `json:"userID"`
	Type     string `json:"type"`
	Reason   string `json:"reason"`
	UserName string `json:""`
}

type ParticipationRewardRequest struct {
	ConferenceID   string `json:"conferenceID"`
	ConferenceType string `json:"conferenceType"`
	Minutes        int32  `json:"minutes"`
}

type UserRewardHistoryRequest struct {
	Sort string `json:"Sort"`
}

type CreateWalletRequest struct {
	UserID string `json:"userID"`
}

type ExclusiveContentRequest struct {
	VideoId string `json:"videoId"`
}
