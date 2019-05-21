package main

import (
	"context"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	_authuc "github.com/yeqown/gRBAC-server/internal-modules/auth/usecase"
	"github.com/yeqown/gRBAC-server/pkg/logger"
)

const (
	// RPCmethodName is inner name to easily understand
	RPCmethodName = "Auth.IsPermitted"
)

// StartRPC ...
func StartRPC(ctx context.Context, port int) {
	// init server and struct
	srv := rpc.NewServer()
	authRPC := &_authuc.Auth{}

	// initial listener
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		panic(err)
	}
	logger.Logger.Infof("rpc listening on: %v", port)
	defer listener.Close()

	// register
	srv.Register(authRPC)

	// loop accept and serve conn
	for {
		select {
		case <-ctx.Done():
			logger.Logger.Info("rpc server quit")
			return
		default:
			conn, err := listener.Accept()
			if err != nil {
				logger.Logger.Errorf("could not accept connection: %v", err)
				continue
			}
			// handler connection
			go srv.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}
}
