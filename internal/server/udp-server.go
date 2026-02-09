package server

import (
	"fmt"
	"log"
	"net"
)

func UDPServer() {
	addr, err := net.ResolveUDPAddr("udp", ":9001")
	if err != nil {
		log.Printf("udp resolve addr error: %v", err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Printf("udp listen error: %v", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)

	fmt.Printf("UDP Server running on port %d\n", addr.Port)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			continue
		}

		msg := string(buf[:n])
		fmt.Println("Received:", msg)

		_, err = conn.WriteToUDP(buf[:n], clientAddr)
		if err != nil {
			fmt.Println("Possible packet loss or client unreachable")
		}
	}
}
