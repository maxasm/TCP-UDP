package main

import (
	"log"
	"net"
	"os"
	"fmt"
	"io"
)

// helper function to handle incomming TCP connections
func handle_TCP_connection(conn *net.TCPConn) {
	// Read data from the connection
	data, err__read_data := io.ReadAll(conn)
	if err__read_data != nil {
		log.Printf("[Error]: %s\n", err__read_data)
		// conn.Close()
	}

	fmt.Printf("---- data from TCP connection ----\n%s\n", string(data))
	conn.Close()
}

func main() {
	// Listen for TCP connections on port 8080
	listener, err__create_listener := net.Listen("tcp", "127.0.0.1:8080")
	if err__create_listener != nil {
		log.Printf("[Error]: %s\n", err__create_listener)
		os.Exit(1)
	}

	fmt.Printf("[Event]: started listening for tcp connection on -> %s\n", listener.Addr().String())
	// Listen for incomming connections
	for {
		// Waits for a connection and thus is blocking
		conn, err__make_conn := listener.Accept()
		if err__make_conn != nil {
			log.Printf("[Error]: %s\n", err__make_conn)
		}

		// handle the connection in a different go-routine
		go func(){
			// convert conn into a TCP connection
			if tcp_conn,ok := conn.(*net.TCPConn); ok {
				handle_TCP_connection(tcp_conn)
			} else {
				log.Printf("[Warning]: connection is not a TCP connection\n")
			}
		}()
	}
}
