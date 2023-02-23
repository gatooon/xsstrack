package client

import (
	"net"
	"xsstrack/servers/web"
)

func RunClientServer() {

	*isRequestReceived = false

	SrvInfo := serverInfo{
		comType: "tcp",
		addr:    "localhost:8989",
	}

	for {
		go waitNewClient(SrvInfo)
		if *isRequestReceived == true {
			*isRequestReceived = false
			for client := web.HandledClientList.Front(); client != nil; client = client.Next() {
				clientConn := client.Value.(web.HandledClient)
				sendDataToClient(clientConn.ClientConn, *receivedData)
			}
		}
	}
}

func sendDataToClient(conn net.Conn, receivedData web.RequestDataStruct) {
	conn.Write([]byte(receivedData.Url))
	conn.Write([]byte(receivedData.Header))
	conn.Write([]byte(receivedData.Body))
}
