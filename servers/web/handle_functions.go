package web

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type triggerFunc func(http.ResponseWriter, string, string)

func specialUrl(writer http.ResponseWriter, request *http.Request, envVarName string, function triggerFunc, sendType string) {
	fileName := os.Getenv(envVarName) + request.URL.String()
	_, err := os.Stat(fileName)
	fmt.Println(fileName)
	if os.IsNotExist(err) {
		fmt.Println("payload not sent")
	} else {
		function(writer, fileName, sendType)
		fmt.Println(request.URL.String() + " sent")
	}
}

func catch(writer http.ResponseWriter, request *http.Request) {
	switch {
	case strings.HasPrefix(request.URL.String(), "/payloads/"):
		specialUrl(writer, request, "XSSTRACK_PAYLOADS_FOLDER", sendPayload, "application/javascript")
	case strings.HasPrefix(request.URL.String(), "/web/"):
		specialUrl(writer, request, "XSSTRACK_WEB_FOLDER", sendPayload, "text/html")
	}

	RequestData.Url = request.URL.String()
	RequestData.Header = convertHeaderToString(request.Header)

	body, _ := io.ReadAll(request.Body)
	RequestData.Body = string(body)

	IsRequestReceived = true

	fmt.Println(request.URL.String())
}

func sendPayload(writer http.ResponseWriter, fileName string, sendType string) {

	data, _ := os.Open(fileName)

	dataInfo, _ := data.Stat()

	dataSize := dataInfo.Size()

	buffer := make([]byte, dataSize)
	data.Read(buffer)

	writer.Header().Set("Content-Type", sendType)
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Write(buffer)

}
