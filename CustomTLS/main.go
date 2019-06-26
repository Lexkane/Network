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
		go logSNI(conn)
	}

}

func logSNI(conn net.Conn) {
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	var buf bytes.Buffer
	if _, err := io.CopyN(&buf, conn, 1+2+2); err != nil {
		log.Println(err)
		return
	}
	length := binary.BigEndian.Uint16(buf.Bytes()[3:5])

	if _, err := io.CopyN(&buf, conn, int64(length)); err != nil {
		log.Println(err)
		return
	}
	//Special Parser based on cryptobytes
	ch, ok := ParseClientHello(buf.Bytes())
	if ok {
		log.Printf("Got a connection with SNI %q", ch.SNI)
	}
	c := perfixConn{
		Conn:   conn,
		Reader: io.MultiReader(&buf, conn),
	}
	conn.SetReadDeadline(time.Time{})
	cert, err := tls.LoadX509KeyPair("localhost.pem", "localhost-key.pem")
	if err != nil {
		log.Fatal(err)
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	tlsConn := tls.Server(c, config)
	copyToStderr(tlsconn)

}

type prefixConn struct {
	nett.Conn
	io.Reader
}

func (c prefixConn) Read(p []byte) (int, error) {
	return c.Reader.Read(p)
}

func proxy(conn net.Conn) {
	defer conn.Close()
	conn, err := net.Dial("tcp", "gophercon.com:443")
	if err != nil {
		log.Println(err)
		return
	}
	defer remote.Close()

	go io.Copy(remote, conn)
	io.Copy(conn, remote)
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
