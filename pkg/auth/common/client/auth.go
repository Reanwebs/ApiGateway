package client

import (
	"fmt"
	"gateway/pkg/auth/common"
	"gateway/pkg/auth/common/pb"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AutharizationClient
}

func InitServiceClient(c *common.Config) pb.AutharizationClient {

	cc, err := grpc.Dial(c.Port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAutharizationClient(cc)
}
