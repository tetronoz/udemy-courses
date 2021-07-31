package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error listening on port 8080 %v", err)
	}

	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("Error accepting connection %v", err)	
		}

		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if line == "" {
			break
		}
	}

	io.WriteString(conn, "I see you")
}