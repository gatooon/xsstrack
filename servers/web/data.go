package web

import (
	"container/list"
	"encoding/json"
	"net"
)

type RequestDataStruct struct {
	Url    string
	Header string
	Body   string
}

type HandledClient struct {
	IsClientInfoReceived bool
	ClientConn           net.Conn
}

var RequestData RequestDataStruct

var HandledClientList list.List
var IsRequestReceived bool
var documentType string

func convertHeaderToString(inputMap map[string][]string) (outputString string) {
	data, _ := json.Marshal(inputMap)
	outputString = string(data)
	return
}
