package main

import (
	"flag"

	"github.com/yeqown/gRBAC-server/logger"
)

var (
	logpath = flag.String("logpath", "./logs", "save log files in this folders, default `./logs`")
	port    = flag.Int("port", 8080, "http server port, default is 8080")
	rpcPort = flag.Int("rport", 8081, "rpc server listen port, default is 8081")
)

func main() {
	flag.Parse()

	// init logger
	logger.InitLogger(*logpath)

	// release mode
	// gin.SetMode(gin.ReleaseMode)

	// start HTTP server
	StartHTTP(*port)

	// start RPC server
	// only provide permission check function ~
	StartRPC(*rpcPort)
}
