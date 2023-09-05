package interfaces

import (
	"context"
	"gateway/pkg/common/pb"
	"mime/multipart"
)

type VideoClient interface {
	UploadVideo(context.Context, *multipart.FileHeader) (*pb.UploadVideoResponse, error)
	StreamVideo(context.Context, string, string) (pb.VideoService_StreamVideoClient, error)
	FindAllVideo(context.Context) (*pb.FindAllResponse, error)
}
