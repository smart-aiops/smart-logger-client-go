package main

import (
	"fmt"
	"os"
	"os/signal"
	"smart/logger/client/demo/grpctrans"
	"smart/logger/client/demo/pkg/httpx"
	"smart/logger/client/demo/router"
	"syscall"
)

func main() {
	code := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cleanFunc, err := initialize()
	if err != nil {
		fmt.Printf("server init fail: %s", err)
		os.Exit(code)
	}

EXIT:
	for {
		sig := <-sc
		fmt.Printf("received signal: %s", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			code = 0
			break EXIT
		case syscall.SIGHUP:
			// reload configuration?
		default:
			break EXIT
		}
	}

	cleanFunc()
	fmt.Printf("server exited")
	os.Exit(code)

}

func initialize() (func(), error) {

	fns := Functions{}

	cleanGrpcConn, err := grpctrans.Run("127.0.0.1", 8000)
	if err != nil {
		fmt.Printf("grpctrans start failed, err: %v", err)
		return fns.Ret(), err
	} else {
		fns.Add(cleanGrpcConn)
	}

	// init http server
	r := router.InitRouter()
	httpClean := httpx.InitServer(8888, r)
	fns.Add(httpClean)

	// release all the resources
	return fns.Ret(), nil
}

type Functions struct {
	List []func()
}

func (fs *Functions) Add(f func()) {
	fs.List = append(fs.List, f)
}

func (fs *Functions) Ret() func() {
	return func() {
		for i := 0; i < len(fs.List); i++ {
			fs.List[i]()
		}
	}
}
