package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return //
		}

		time.Sleep(1 * time.Second)
	}
}

/*
What does this have anything to do with go routines?

Turn out it doesn't. This go script creates a TCP server that listen on port 8080, this server only accept 1 client at a time by the `handleConn` function.
*/
