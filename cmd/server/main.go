package main

import (
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatalf("Start server error: %v", err)
		return
	}
	defer server.Close()
	log.Printf("Server started on %s, waiting for connection.", ":6379")
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalf("Accept error: %v", err)
			return
		}
		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		log.Printf("Read: %s", buf)
		if err != nil {
			if err.Error() == "EOF" {
				log.Printf("Connection closed by client.")
				return
			}
			log.Fatalf("Read error: %v", err)
			return
		}
		conn.Write([]byte("+OK\r\n"))
	}
}
