package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"tcp/util"
)

func main() {
	fmt.Println("proxy listening for connection@", util.ProxyAddress)
	// listen on all interfaces
	ln, _ := net.Listen("tcp", util.ProxyAddress)
	// accept connection on port
	proxyConn, _ := ln.Accept()
	defer proxyConn.Close()
	// connect to proxy server
	serverConn, err := net.Dial("tcp", util.ServerAddress)
	if err != nil {
		os.Exit(1)
	}
	defer proxyConn.Close()
	receiveFile(proxyConn, serverConn)
}

func receiveFile(conn, proxyConn net.Conn) {
	_, err := io.Copy(proxyConn, conn)
	if err != nil {
		log.Fatal(err)
	}
}
