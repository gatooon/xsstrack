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

func waitNewClient(srvInfo serverInfo) {
	for {
		srv, _ := net.Listen(srvInfo.comType, srvInfo.addr)
		fmt.Println("Server UP, waiting for connexion")

		conn, _ := srv.Accept()

		buffer := make([]byte, 1024)
		mLen, _ := conn.Read(buffer)

		web.HandledClientList.PushBack(
			web.HandledClient{
				IsClientInfoReceived: true,
				UrlListening:         string(buffer[:mLen]),
				ClientConn:           conn,
			})

		fmt.Println("Client " + conn.RemoteAddr().String() + " Connected")
	}
}
