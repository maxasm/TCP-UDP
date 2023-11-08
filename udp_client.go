package main

import (
	"log"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err__make_udp_conn := net.Dial("udp4", "127.0.0.1:7070")
	if err__make_udp_conn != nil {
		log.Printf("[Error]: %s\n", err__make_udp_conn)
		os.Exit(1)
	}

	fmt.Printf("[Event]: made a udp connection to -> %s\n", conn.RemoteAddr().String())
	// write a test message and send it via UDP
	conn.Write([]byte("Hello Test"))
	conn.Close()
}
