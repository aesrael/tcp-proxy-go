package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"tcp/util"
)

const filename = "file.txt"

func main() {
	fmt.Println("server listening @", util.ServerAddress)
	// listen on all interfaces
	ln, err := net.Listen("tcp", util.ServerAddress)
	util.HandleErr(err)

	//accept connection on port
	conn, err := ln.Accept()
	util.HandleErr(err)
	receiveFile(conn)
}

func receiveFile(conn net.Conn) {
	// create new file
	file, err := os.Create(filename)
	util.HandleErr(err)
	fmt.Println("server received file")
	defer file.Close()
	// accept data from client & write to a new file
	_, err = io.Copy(file, conn)
	util.HandleErr(err)
}
