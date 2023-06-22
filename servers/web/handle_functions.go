package web

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type triggerFunc func(http.ResponseWriter, string, string)

func specialUrl(writer http.ResponseWriter, request *http.Request, envVarName string, function triggerFunc) {
	fileName := strings.TrimSuffix(os.Getenv(envVarName), "/") + request.URL.String()
	_, err := os.Stat(fileName)
	fileExtension := filepath.Ext(fileName)
	var mime string
	switch fileExtension {
	case ".js":
		mime = "application/javascript"
	case ".css":
		mime = "text/css"
	default:
		mime = "text/html"
	}

	fmt.Println(fileName)
	if os.IsNotExist(err) {
		fmt.Println("payload not sent")
	} else {
		function(writer, fileName, mime)
		fmt.Println(request.URL.String() + " sent")
	}
}

func catch(writer http.ResponseWriter, request *http.Request) {
	switch {
	case strings.HasPrefix(request.URL.String(), "/payloads/"):
		specialUrl(writer, request, "XSSTRACK_PAYLOADS_FOLDER", sendPayload)
	case strings.HasPrefix(request.URL.String(), "/web/"):
		specialUrl(writer, request, "XSSTRACK_WEB_FOLDER", sendPayload)
	}

	RequestData.Url = request.URL.String()
	RequestData.Header = convertHeaderToString(request.Header)

	body, _ := io.ReadAll(request.Body)
	RequestData.Body = string(body)

	IsRequestReceived = true

	fmt.Println(request.URL.String())
}

func sendPayload(writer http.ResponseWriter, fileName string, mime string) {

	data, _ := os.Open(fileName)

	dataInfo, _ := data.Stat()

	dataSize := dataInfo.Size()

	buffer := make([]byte, dataSize)
	data.Read(buffer)

	writer.Header().Set("Content-Type", mime)
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Write(buffer)

}
