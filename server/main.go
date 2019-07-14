package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("connecting")
	// listen on all interfaces
	ln, _ := net.Listen("tcp", "127.0.0.1:8081")
	// accept connection on port
	conn, _ := ln.Accept()
	defer conn.Close()
	const filename = "file.txt"
	receiveFile(filename, conn)
}

func receiveFile(filename string, conn net.Conn) {
	// accept connection
	// create new file
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// accept file from client & write to a new file
	_, err = io.Copy(file, conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file succesfully created!")
}
