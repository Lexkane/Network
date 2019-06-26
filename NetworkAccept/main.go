package main

import (
	"io"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "localhostL:4242")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go copyToStderr(conn)
	}

}

func copyToStderr(conn net.Conn) {
	n, err := io.Copy(os.Stderr, conn)
	log.Printf("Copied %d bytes; finished with err -%v", n, err)
}
