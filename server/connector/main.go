package main

import "github.com/GoodOneGuy/jiochat/server/connector/server"

func main() {

	// 启动websocket服务
	if err := server.StartWBSocket(); err != nil {

		return
	}

}
