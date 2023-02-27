package web

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func catch(writer http.ResponseWriter, request *http.Request) {
	if strings.HasPrefix(request.URL.String(), "/payloads/") {
		fileName := "." + request.URL.String()
		_, err := os.Stat(fileName)
		if os.IsNotExist(err) {
			fmt.Println("payload not sent")
		} else {
			sendPayload(writer, fileName)
			fmt.Println(request.URL.String() + " sent")
		}
	}
	RequestData.Url = request.URL.String()
	RequestData.Header = convertHeaderToString(request.Header)

	body, _ := io.ReadAll(request.Body)
	RequestData.Body = string(body)

	IsRequestReceived = true

	fmt.Println(request.URL.String())
}

func sendPayload(writer http.ResponseWriter, fileName string) {

	data, _ := os.Open(fileName)

	buffer := make([]byte, 1024)
	data.Read(buffer)

	writer.Header().Set("Content-Type", "text/javascript")
	writer.Write(buffer)

}
