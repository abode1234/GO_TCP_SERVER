package main

import (
	"tcpserver/cmd/server"
	"tcpserver/cmd/tcp"
)



func main( ) {
  server.Server()
  tcp.TcpServer()

}
