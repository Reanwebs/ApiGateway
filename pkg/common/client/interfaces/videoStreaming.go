package interfaces

import (
	"context"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb/video"
	"mime/multipart"
)

type VideoClient interface {
	UploadVideo(context.Context, *multipart.FileHeader, models.UploadVideo) (*video.UploadVideoResponse, error)
	FetchVideos(context.Context, models.FetchVideosRequest) (*video.FindUserVideoResponse, error)
	FetchArchivedVideos(context.Context, models.FetchVideosRequest) (*video.FindArchivedVideoByUserIdResponse, error)
	FetchAllVideos(context.Context) (*video.FetchAllVideoResponse, error)
	ArchiveVideo(ctx context.Context, request models.ArchivedVideos) (*video.ArchiveVideoResponse, error)
	GetVideoById(context.Context, string) (*video.GetVideoByIdResponse, error)
}
