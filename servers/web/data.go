package web

import (
	"encoding/json"
)

type RequestDataStruct struct {
	Url    string
	Header string
	Body   string
}

var IsRequestReceived bool

var RequestData RequestDataStruct

func convertHeaderToString(inputMap map[string][]string) (outputString string) {
	data, _ := json.Marshal(inputMap)
	outputString = string(data)
	return
}
