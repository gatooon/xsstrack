package client

import (
	"fmt"
	"net"
	"xsstrack/servers/web"
)

func RunClientServer() {

	*isRequestReceived = false

	SrvInfo := serverInfo{
		comType: "tcp",
		addr:    "localhost:8800",
	}

	srv, err := net.Listen(SrvInfo.comType, SrvInfo.addr)
	if err != nil {
		fmt.Println("Error : ", err)
	}
	fmt.Println("Server UP, waiting for connexion")
	go waitNewClient(srv)
	for {
		if *isRequestReceived == true {
			*isRequestReceived = false
			for client := web.HandledClientList.Front(); client != nil; client = client.Next() {
				clientConn := client.Value.(web.HandledClient)
				if "/"+clientConn.UrlListening == receivedData.Url {
					IsClientAlive := sendDataToClient(clientConn.ClientConn, *receivedData)
					if IsClientAlive == false {
						web.HandledClientList.Remove(client)
					}
				}
			}
		}
	}
}

func sendDataToClient(conn net.Conn, receivedData web.RequestDataStruct) (IsClientAlive bool) {
	_, err := conn.Write([]byte(receivedData.Url + "\n" + receivedData.Header + "\n" + receivedData.Body))
	IsClientAlive = true
	if err != nil {
		IsClientAlive = false
	}
	return
}
