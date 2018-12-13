package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/yeqown/gRBAC-server/logger"
	"github.com/yeqown/gRBAC-server/services"
)

const (
	// RPCmethodName is inner name to easily understand
	RPCmethodName = "Auth.IsPermitted"
)

// StartRPC ...
func StartRPC(port int) {
	// init server and struct
	srv := rpc.NewServer()
	authRPC := &services.Auth{}

	// initial listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	// register
	srv.Register(authRPC)

	// loop accept and serve conn
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Logger.Errorf("could not accept connection: %v", err)
			continue
		}
		// handler connection
		go srv.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
