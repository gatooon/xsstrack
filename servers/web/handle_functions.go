package web

import (
	"fmt"
	"io"
	"net/http"
)

func catch(_ http.ResponseWriter, request *http.Request) {

	RequestData.Url = request.URL.String()
	RequestData.Header = convertHeaderToString(request.Header)

	body, _ := io.ReadAll(request.Body)
	RequestData.Body = string(body)

	IsRequestReceived = true

	fmt.Println(request.URL.String())
}
