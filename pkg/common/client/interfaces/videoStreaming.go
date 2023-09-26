package interfaces

import (
	"context"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb/video"
	"mime/multipart"
)

type VideoClient interface {
	UploadVideo(context.Context, *multipart.FileHeader, models.UploadVideo) (*video.UploadVideoResponse, error)
}
