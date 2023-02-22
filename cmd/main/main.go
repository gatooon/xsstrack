package main

import (
	"xsstrack/servers/client"
	"xsstrack/servers/web"
)

func main() {
	go client.RunClientServer()
	web.RunHttpServer()
}
