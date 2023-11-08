// This is a simple implementation of a UDP server
package main

import (
	"sync"
	"net"
	"log"
	"fmt"
	"os"
)

// TODO: implement
func read_all_from_udp(conn net.UDPConn) {
	
}

func handle_UDP_connection(conn *net.UDPConn) {
	// read a packet from the UDP connection
	buffer := make([]byte, 512)
	// read from the UDP connection
	// TODO: Read from the UDP connection until you get an error that there is no more data
	b_count,_, err__read_from_udp := conn.ReadFrom(buffer)
	if err__read_from_udp != nil {
		log.Printf("[Error]: %s\n", err__read_from_udp)
		conn.Close()
	}

	fmt.Printf("[Event]: read %d bytes from the UDP connection\n", b_count)
	fmt.Printf("data -> %s\n",string(buffer))
}

func main() {
	// Listen on port 7070 for UDP connections
	udp_conn, err__make_udp_conn := net.ListenUDP("udp4", &net.UDPAddr{
		Port: 7070,
	})

	if err__make_udp_conn != nil {
		log.Printf("[Error]: %s\n", err__make_udp_conn)
		os.Exit(1)
	}

	log.Printf("[Event]: listening for UDP connections on -> %s\n", udp_conn.LocalAddr().String())

	// create a wait group to manage the go-routines
	var wg sync.WaitGroup

	// TODO: You actually do not need a WaitGroup and a different Go Routine
	wg.Add(1)
	go func(wg sync.WaitGroup){
		for {
			// handle the connection in a separate Go Routine
			handle_UDP_connection(udp_conn)
		}
		wg.Done()
	}(wg)

	wg.Wait()
}
