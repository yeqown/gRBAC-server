package main

import (
	"context"
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/yeqown/gRBAC-server/logger"
	"github.com/yeqown/gRBAC-server/services"
)

var (
	logpath = flag.String("logpath", "./logs", "save log files in this folders, default `./logs`")
	port    = flag.Int("port", 8080, "http server port, default is 8080")
	rpcPort = flag.Int("rport", 8081, "rpc server listen port, default is 8081")
	openRPC = flag.Bool("openrpc", true, "open or not open rpc server")
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flag.Parse()

	// init logger
	logger.InitLogger(*logpath)

	// init secret file and token
	services.InitSecretFile()

	// set gin release mode
	gin.SetMode(gin.ReleaseMode)

	if *openRPC {
		// start RPC server
		// only provide permission check function ~
		go StartRPC(ctx, *rpcPort)
	}

	// start HTTP server
	StartHTTP(*port)
}
