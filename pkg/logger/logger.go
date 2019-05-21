package logger

import (
	log "github.com/yeqown/infrastructure/framework/logrus-logger"
)

var (
	// Logger ...
	Logger *log.Logger
)

// InitLogger ...
func InitLogger(logpath string) (err error) {
	Logger, err = log.NewJSONLogger(logpath, "server.log", "debug")
	return
}
