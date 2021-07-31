package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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
	lineNumber := 0
	var Method, URI string

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if lineNumber == 0 {
			splittedLine := strings.Split(line, " ")
			Method = splittedLine[0]
			URI = splittedLine[1]
		}
		lineNumber++
		if line == "" {
			break
		}
	}

	body := "I see you connected"
	body += "\n"
	body += "Using method: " + Method
	body += "\n"
	body += "and URI: " + URI
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}