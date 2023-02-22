package client

import "xsstrack/servers/web"

type serverInfo struct {
	comType string
	addr    string
}

var receivedData = &web.RequestData
var isRequestReceived = &web.IsRequestReceived
