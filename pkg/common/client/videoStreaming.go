package client

import (
	"context"
	"errors"
	"fmt"
	"gateway/pkg/common/client/interfaces"
	"gateway/pkg/common/config"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb/video"
	"io"
	"mime/multipart"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type videoClient struct {
	Server video.VideoServiceClient
}

func InitVideoStreamingClient(c *config.Config) (interfaces.VideoClient, error) {
	cc, err := grpc.Dial(c.VideoService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return NewVideoClient(video.NewVideoServiceClient(cc)), nil
}

func NewVideoClient(server video.VideoServiceClient) interfaces.VideoClient {
	return &videoClient{
		Server: server,
	}
}

func (c *videoClient) UploadVideo(ctx context.Context, file *multipart.FileHeader, request models.UploadVideo) (*video.UploadVideoResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}

	upLoadfile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer upLoadfile.Close()

	stream, err := c.Server.UploadVideo(ctx)
	if err != nil {
		return nil, errors.New("streaming service down")
	}
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
			UserId:      userId,
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

func (c *videoClient) FetchVideos(ctx context.Context, request models.FetchVideosRequest) (*video.FindUserVideoResponse, error) {
	res, err := c.Server.FindUserVideo(ctx, &video.FindUserVideoRequest{
		UserName: request.UserName,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *videoClient) FetchArchivedVideos(ctx context.Context, request models.FetchVideosRequest) (*video.FindArchivedVideoByUserIdResponse, error) {
	res, err := c.Server.FindArchivedVideoByUserId(ctx, &video.FindArchivedVideoByUserIdRequest{
		UserName: request.UserName,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *videoClient) FetchAllVideos(ctx context.Context) (*video.FetchAllVideoResponse, error) {
	res, err := c.Server.FetchAllVideo(ctx, &video.FetchAllVideoRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *videoClient) ArchiveVideo(ctx context.Context, request models.ArchivedVideos) (*video.ArchiveVideoResponse, error) {

	videoID, err := strconv.ParseUint(request.VideoId, 10, 32)
	if err != nil {
		fmt.Println("Error converting video ID:", err)
		return nil, errors.New("video id mismatching")
	}
	res, err := c.Server.ArchiveVideo(ctx, &video.ArchiveVideoRequest{
		VideoId: uint32(videoID),
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *videoClient) GetVideoById(ctx context.Context, request models.GetVideoById) (*video.GetVideoByIdResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		fmt.Println("userId not found in context.")
		return nil, errors.New("login again")
	}
	res, err := c.Server.GetVideoById(ctx, &video.GetVideoByIdRequest{
		VideoId:  request.VideoId,
		UserName: request.UserNAme,
		UserId:   userId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *videoClient) ToggleStar(ctx context.Context, request models.ToggleStarRequest) (*video.ToggleStarResponse, error) {

	res, err := c.Server.ToggleStar(ctx, &video.ToggleStarRequest{
		VideoId:  request.VideoId,
		UserNAme: request.UserNAme,
		Starred:  request.Starred,
	})
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (c *videoClient) BlockVideo(ctx context.Context, request models.BlockVideoRequest) (*video.BlockVideoResponse, error) {

	res, err := c.Server.BlockVideo(ctx, &video.BlockVideoRequest{
		VideoId: request.VideoId,
		Reason:  request.Reason,
		Block:   request.Block,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *videoClient) ReportVideo(ctx context.Context, request models.ReportVideoRequest) (*video.ReportVideoResponse, error) {
	res, err := c.Server.ReportVideo(ctx, &video.ReportVideoRequest{
		VideoId: request.VideoId,
		Reason:  request.Reason,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *videoClient) GetReportedVideos(ctx context.Context) (*video.GetReportedVideosResponse, error) {

	res, err := c.Server.GetReportedVideos(ctx, &video.GetReportedVideosRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}
