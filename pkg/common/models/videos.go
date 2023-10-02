package models

type UploadVideo struct {
	UserName    string `json:"userName"`
	AvatarId    string `json:"avatarId"`
	Title       string `json:"title"`
	Discription string `json:"discription"`
	Interest    string `json:"interest"`
	ThumbnailId string `json:"thumbnailId"`
}

type FetchVideosRequest struct {
	UserName string `json:"userName"`
}

type ArchivedVideos struct {
	VideoId string `json:"videoId"`
}

type ToggleStarRequest struct {
	VideoId  string `json:"videoId"`
	UserNAme string `json:"userName"`
	Starred  bool   `json:"starred"`
}

type GetVideoById struct {
	VideoId  string
	UserNAme string
}

type BlockVideoRequest struct {
	VideoId string `json:"videoId"`
	Reason  string `json:"reason"`
	Block   bool   `json:"block"`
}

type ReportVideoRequest struct {
	VideoId string `json:"videoId"`
	Reason  string `json:"reason"`
}
