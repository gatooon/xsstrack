# xsstrack

Server that receive request and forward it to a client.<br>
Can also send html and javascript files.

# Info
Web port is 80 ( determined in servers/web/web_server.go )<br>
Client port is 8800 ( determined in servers/client/client_server.go )

# Requirement 

golang

# Installation/Configuration

clone the repository<br>
git clone https://github.com/gatooon/xsstrack.git

Build project<br>
GOOS=linux GOARCH=amd64 go build -o xsstrack cmd/main/main.go

## Optional 
Create 2 environment variables with value of the payloads folder and the web folder<br>
folder have to be named "web" and "payloads"

export XSSTRACK_PAYLOADS_FOLDER<br>
export XSSTRACK_WEB_FOLDER<br>

exemples:<br>
/home/ping/web/ -> export XSSTRACK_WEB_FOLDER="/home/ping/"<br>
/home/pong/payloads/ -> export XSSTRACK_PAYLODS_FOLDER="/home/pong/"

# Run client
Clone and build the following project:<br>
https://github.com/gatooon/xsstrack-client
