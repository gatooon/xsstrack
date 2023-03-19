package client

import (
	"fmt"
	"net"
	"xsstrack/servers/web"
)

type serverInfo struct {
	comType string
	addr    string
}

var receivedData = &web.RequestData
var isRequestReceived = &web.IsRequestReceived

func waitNewClient(srv net.Listener) {
	for {
		conn, _ := srv.Accept()
		web.HandledClientList.PushBack(
			web.HandledClient{
				IsClientInfoReceived: true,
				ClientConn:           conn,
			})
		fmt.Println("Client " + conn.RemoteAddr().String() + " Connected")
	}
}
