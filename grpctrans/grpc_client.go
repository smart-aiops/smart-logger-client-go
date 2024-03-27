package grpctrans

import (
	"fmt"
	"github.com/smart-aiops/smart-logger-client-go/client"
)

var GrpcClientConf *client.GrpcClient

func Run(ip string, port int) (func(), error) {
	var err error
	GrpcClientConf, err = client.NewGrpcClient(ip, port)
	if err != nil {
		fmt.Printf("grpc client start failed, err:%v", err)
	}

	return GrpcClientConf.Destroy(), nil
}
