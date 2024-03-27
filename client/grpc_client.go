package client

import (
	"errors"
	"fmt"
	"github.com/smart-aiops/smart-logger-client-go/proto"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	client proto.TransferServiceClient
	conn   *grpc.ClientConn
}

func NewGrpcClient(ip string, port int) (*GrpcClient, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())
	if err != nil {
		return nil, errors.New("connect gPRC server failed")
	}

	grpcClient := &GrpcClient{
		client: proto.NewTransferServiceClient(conn),
		conn:   conn,
	}

	return grpcClient, nil
}

func (g *GrpcClient) transfer(data proto.RawData) {
	g.transfer(data)
}

func (g *GrpcClient) destroy() func() {
	return func() {
		g.conn.Close()
	}
}
