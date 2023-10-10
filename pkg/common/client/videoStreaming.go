package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"gateway/pkg/common/client/interfaces"
	"gateway/pkg/common/config"
	"gateway/pkg/common/models"
	"gateway/pkg/common/pb/video"
	"io"
	"log"
	"mime/multipart"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type videoClient struct {
	Server video.VideoServiceClient
}

func InitVideoStreamingClient(c *config.Config) (interfaces.VideoClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	caCert, err := os.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, fmt.Errorf("failed to read CA certificate: %w", err)
	}

	// create cert pool and append ca's cert
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		return nil, errors.New("failed to append CA certificate to certificate pool")
	}

	//read client cert
	clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if err != nil {
		return nil, fmt.Errorf("failed to load client certificate and key: %w", err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	tlsCredential := credentials.NewTLS(config)

	conn, err := grpc.DialContext(ctx, c.VideoService, grpc.WithTransportCredentials(tlsCredential))
	if err != nil {
		return nil, fmt.Errorf("failed to load client certificate and key: %w", err)
	}

	return NewVideoClient(video.NewVideoServiceClient(conn)), nil
}

func NewVideoClient(server video.VideoServiceClient) interfaces.VideoClient {
	return &videoClient{
		Server: server,
	}
}

func (c *videoClient) UploadVideo(ctx context.Context, file *multipart.FileHeader, request models.UploadVideo) (*video.UploadVideoResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		log.Println("userId not found in context.")
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
			UserName:     request.UserName,
			AvatarId:     request.AvatarId,
			Intrest:      request.Interest,
			ThumbnailId:  request.ThumbnailId,
			Title:        request.Title,
			Discription:  request.Discription,
			Filename:     file.Filename,
			UserId:       userId,
			Data:         buffer[:n],
			Exclusive:    request.Exclusive,
			CoinForWatch: request.Coin_for_watch,
		}); err != nil {
			log.Println("error in streaming in client", err)
			return nil, err
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Println("error in response ", err)
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

	res, err := c.Server.ArchiveVideo(ctx, &video.ArchiveVideoRequest{
		VideoId: request.VideoId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *videoClient) GetVideoById(ctx context.Context, request models.GetVideoById) (*video.GetVideoByIdResponse, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		log.Println("userId not found in context.")
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

func (c *videoClient) FetchExclusiveVideo(ctx context.Context) (*video.FetchExclusiveVideoResponse, error) {

	res, err := c.Server.FetchExclusiveVideo(ctx, &video.FetchExclusiveVideoRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}
