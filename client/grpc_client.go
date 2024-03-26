package client

import (
	"fmt"
	"github.com/cooper190113/smart-logger-client-go/proto"
	"github.com/toolkits/pkg/logger"
	"google.golang.org/grpc"
)

var GrpcClient proto.TransferServiceClient

// Init grpc server ip and port
func Init(ip string, port int) (func(), error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())
	if err != nil {
		logger.Errorf("connect gPRC server failed, err: %v", err)
	}

	GrpcClient = proto.NewTransferServiceClient(conn)
	return func() {
		logger.Errorf("GrpcClient exiting")
		conn.Close()
	}, nil
}
