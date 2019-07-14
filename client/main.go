package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const address = "127.0.0.1:8081"

func main() {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()
	const filename = "./file.txt"
	createFile(filename)
	sendFile(filename, conn)
}

func sendFile(filename string, conn net.Conn) {
	// open file to send to server
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// send file
	_, err = io.Copy(conn, file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file sent!")
}

func createFile(filename string) {
	const content = "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor"

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		for i := 1; i <= 1e3; i++ {
			file.WriteString(content)
		}
	}
}
