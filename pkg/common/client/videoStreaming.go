package client

import (
	"context"
	"fmt"
	"gateway/pkg/common/client/interfaces"
	"gateway/pkg/common/config"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb/video"
	"io"
	"mime/multipart"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type videoClient struct {
	Server video.VideoServiceClient
}

func InitVideoStreamingClient(c *config.Config) (interfaces.VideoClient, error) {
	creds, err := credentials.NewClientTLSFromFile("server.crt", "")
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(c.VideoService, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	return NewVideoClient(video.NewVideoServiceClient(conn)), nil
}

func NewVideoClient(server video.VideoServiceClient) interfaces.VideoClient {
	return &videoClient{
		Server: server,
	}
}

func (c *videoClient) UploadVideo(ctx context.Context, file *multipart.FileHeader, request models.UploadVideo) (*video.UploadVideoResponse, error) {
	upLoadfile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer upLoadfile.Close()

	stream, _ := c.Server.UploadVideo(ctx)
	chunkSize := 4096
	buffer := make([]byte, chunkSize)

	for {
		n, err := upLoadfile.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if err := stream.Send(&video.UploadVideoRequest{
			UserName:    request.UserName,
			AvatarId:    request.AvatarId,
			Intrest:     request.Interest,
			ThumbnailId: request.ThumbnailId,
			Title:       request.Title,
			Discription: request.Discription,
			Filename:    file.Filename,
			Data:        buffer[:n],
		}); err != nil {
			fmt.Println("error in streaming in client", err)
			return nil, err
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println("error in response ", err)
		return nil, err
	}

	return response, nil
}
