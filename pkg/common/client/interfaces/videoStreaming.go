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
	ArchiveVideo(ctx context.Context, request models.ArchivedVideos) (*video.ArchiveVideoResponse, error)
}
