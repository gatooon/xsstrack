package client

import (
	"fmt"
	"net"
	"xsstrack/servers/web"
)

func RunClientServer() {

	*isRequestReceived = false

	srvInfo := serverInfo{
		comType: "tcp",
		addr:    "localhost:8800",
	}

	srv, _ := net.Listen(srvInfo.comType, srvInfo.addr)

	fmt.Println("Server UP, waiting for connexion")

	conn, _ := srv.Accept()
	fmt.Println("Client Connected")

	for {
		if *isRequestReceived == true {
			sendDataToClient(conn, *receivedData)
			*isRequestReceived = false
		}
	}
}

func sendDataToClient(conn net.Conn, receivedData web.RequestDataStruct) {
	_, _ = conn.Write([]byte(receivedData.Url))
	_, _ = conn.Write([]byte(receivedData.Header))
	_, _ = conn.Write([]byte(receivedData.Body))
}
