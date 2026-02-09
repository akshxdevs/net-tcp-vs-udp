package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync/atomic"
)

var connCount int64

func handleConn(c net.Conn) {
	defer c.Close()
	atomic.AddInt64(&connCount, 1)
	defer atomic.AddInt64(&connCount, -1)

	reader := bufio.NewReader(c)
	name, userID := "unknown", "unknown"
	line, err := reader.ReadString('\n')
	if err == nil {
		parts := strings.Fields(strings.TrimSpace(line))
		if len(parts) >= 3 && strings.EqualFold(parts[0], "HELLO") {
			name = parts[1]
			userID = parts[2]
		}
	}

	fmt.Printf("New connection from %s (%s, id=%s). Active: %d\n",
		c.RemoteAddr().String(),
		name,
		userID,
		atomic.LoadInt64(&connCount),
	)

	for {
		data, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println("Connection Closed:", err)
			return
		}
		c.Write(data)
	}

}

func TCPServer() {
	addr := 9000
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", addr))
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	fmt.Printf("TCP Server running on port %d\n", addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		go handleConn(conn)
	}
}
