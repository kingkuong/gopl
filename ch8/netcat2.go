package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	go mustCopy(os.Stdout, conn) // difference between netcat1 and netcat2, this is ran parallel, acceping input
	mustCopy(conn, os.Stdin)     // difference between netcat 1 and netcat2, this is ran sequential, producing output
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
