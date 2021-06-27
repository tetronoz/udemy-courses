package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer lis.Close()

	for {
		conn, err:= lis.Accept()
		if err != nil {
			log.Println(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	if err := conn.SetDeadline(time.Now().Add(10 * time.Second)); err != nil {
		log.Println("CONN TIMEOUT")
	}

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s", ln)
	}
	defer conn.Close()
}