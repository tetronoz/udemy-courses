package main

import (
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
		con, err := lis.Accept()
		if err != nil {
			log.Fatalf("Error accepting connection %v", err)	
		}

		response := "I see you connected"
		io.WriteString(con, response)
		con.Close()
	}
}