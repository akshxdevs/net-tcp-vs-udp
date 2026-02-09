package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, _ := net.Dial("udp", "localhost:9001")
	defer conn.Close()

	buf := make([]byte, 1024)

	for i := 0; i < 20; i++ {
		msg := fmt.Sprintf("packet-%d", i)
		conn.Write([]byte(msg))

		conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Packet lost:", msg)
		} else {
			fmt.Println("Response:", string(buf[:n]))
		}

		time.Sleep(200 * time.Millisecond)
	}
}
