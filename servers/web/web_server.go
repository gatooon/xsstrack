package web

import (
	"net/http"
)

func RunHttpServer() {
	//for InfoClient.IsClientInfoReceived != true {
	//	time.Sleep(1 * time.Second)
	//}
	//+InfoClient.UrlListening
	http.HandleFunc("/", catch)
	http.ListenAndServe(":8080", nil)
}
