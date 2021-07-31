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

		scanner := bufio.NewScanner(conn)

		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
		defer conn.Close()

		fmt.Println("Code got here.")
		io.WriteString(conn, "I see you connected.")

	}
}