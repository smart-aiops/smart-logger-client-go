package client

import (
	"context"
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

	return &GrpcClient{
		client: proto.NewTransferServiceClient(conn),
		conn:   conn,
	}, nil
}

func (g *GrpcClient) Send(data proto.RawData) (*proto.Response, error) {
	response, err := g.client.Transfer(context.Background(), &data)
	if err != nil {
		fmt.Printf("grpc send data failed, err: %v", err)
		return nil, errors.New("grpc send data failed")
	}
	return response, nil
}

func (g *GrpcClient) Destroy() func() {
	if g.conn != nil {
		return func() {
			g.conn.Close()
		}
	}
	return func() {}
}
