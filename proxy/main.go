package main

import (
	"fmt"
	"good-to-go/utils"
	"io"
	"net"
	"tcp/util"
)

func main() {
	fmt.Println("proxy listening for connection@", util.ProxyAddress)
	// listen on all interfaces
	ln, err := net.Listen("tcp", util.ProxyAddress)
	utils.HandleErr(err)
	// accept connection on port
	proxyConn, err := ln.Accept()
	utils.HandleErr(err)
	defer proxyConn.Close()
	// connect to proxy server
	serverConn, err := net.Dial("tcp", util.ServerAddress)
	utils.HandleErr(err)
	defer proxyConn.Close()
	sendFile(proxyConn, serverConn)
}

func sendFile(conn, proxyConn net.Conn) {
	_, err := io.Copy(proxyConn, conn)
	utils.HandleErr(err)
}
