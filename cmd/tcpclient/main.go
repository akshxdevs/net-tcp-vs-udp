package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr := 9000
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	log.Printf("Tcp Client running on port %d", addr)

	defer conn.Close()

	fmt.Fprintln(conn, "HELLO john 42")

	reader := bufio.NewReader(os.Stdin)

	serverReader := bufio.NewReader(conn)

	for {
		text, _ := reader.ReadString('\n')
		conn.Write([]byte(text))

		resp, _ := serverReader.ReadString('\n')
		fmt.Println("Server: ", resp)
	}
}
