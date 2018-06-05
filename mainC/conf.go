package main

import (
	"bufio"
	"encoding/json"
	"github.com/yeqown/gweb"
	"os"
)

var (
	server_ins *ServerConf
)

// func init() {
// 	server_ins = new(ServerConf)
// }

type ServerConf struct {
	HttpC *gweb.ServerConfig    `json:"HttpServerConfig"`
	RpcC  *gweb.RpcServerConfig `json:"RpcServerConfig"`
	Token string                `json:"token"`
}

func LoadConfig(file string) error {
	fp, err := os.Open(file)
	defer fp.Close()
	if err != nil {
		return err
	}
	body, _ := bufio.NewReader(fp).ReadBytes(0)
	if err = json.Unmarshal(body, &server_ins); err != nil {
		return err
	}
	return nil
}
