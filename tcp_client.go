package main

import (
	"net"
	"log"
)

func main() {
	conn, err__create_conn := net.Dial("tcp", "127.0.0.1:8080")
	if err__create_conn != nil {
		log.Printf("[Error]: %s\n", err__create_conn)
		return
	}

	log.Printf("[Event]: created a tcp connection from: %s\n", conn.LocalAddr())
	_, err__write := conn.Write([]byte("Hello Test"))
	if err__write != nil {
		log.Printf("[Error]: %s\n", err__write)
	}
}
