package web

import (
	"net/http"
)

func RunHttpServer() {
	http.HandleFunc("/", catch)
	http.ListenAndServe(":80", nil)
}
