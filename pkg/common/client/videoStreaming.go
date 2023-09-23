package client

import (
	"context"
	"fmt"
	"gateway/pkg/common/client/interfaces"
	"gateway/pkg/common/config"
	"gateway/pkg/common/pb"
	"io"
	"mime/multipart"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type videoClient struct {
	Server pb.VideoServiceClient
}

func InitVideoStreamingClient(c *config.Config) (interfaces.VideoClient, error) {
	creds, err := credentials.NewClientTLSFromFile("server.crt", "")
	if err != nil {
		return nil, err
	}

	// Dial the gRPC server with TLS
	conn, err := grpc.Dial(c.VideoService, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	return NewVideoClient(pb.NewVideoServiceClient(conn)), nil
}

func NewVideoClient(server pb.VideoServiceClient) interfaces.VideoClient {
	return &videoClient{
		Server: server,
	}
}

func (c *videoClient) UploadVideo(ctx context.Context, file *multipart.FileHeader) (*pb.UploadVideoResponse, error) {
	upLoadfile, err := file.Open()
	if err != nil {
		return nil, err
	}

	fmt.Println("file opned in client")
	defer upLoadfile.Close()
	//getting the stream object
	stream, _ := c.Server.UploadVideo(ctx)
	chunkSize := 4096 // Set your desired chunk size
	buffer := make([]byte, chunkSize)
	for {
		n, err := upLoadfile.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if err := stream.Send(&pb.UploadVideoRequest{
			Filename: file.Filename,
			Data:     buffer[:n],
		}); err != nil {
			fmt.Println("error in streaming in client", err)
			return nil, err
		}
	}
	//the final response is recieved and the streaming is closed
	response, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println("error in response ", err)
		return nil, err
	}

	return response, nil
}

func (c *videoClient) StreamVideo(ctx context.Context, filename, playlist string) (pb.VideoService_StreamVideoClient, error) {
	res, err := c.Server.StreamVideo(ctx, &pb.StreamVideoRequest{
		VideoId: "188e9aed-9629-463a-b76d-a16f9deab6ef",
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *videoClient) FindAllVideo(ctx context.Context) (*pb.FindAllResponse, error) {
	res, err := c.Server.FindAllVideo(ctx, &pb.FindAllRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}
