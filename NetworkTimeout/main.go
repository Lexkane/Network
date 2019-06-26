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
	defer con.Close()
	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Accept))
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Printf("Finished with err=%v", err)
			return
		}
		os.Stderr.Write(buf[:n])
	}
}
