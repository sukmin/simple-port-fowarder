package main

import (
	"flag"
	"io"
	"log"
	"net"
)

func forward(conn net.Conn, clientPort string) {
	client, err := net.Dial("tcp", clientPort)
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	log.Printf("Connected to localhost %v\n", conn)
	go func() {
		defer client.Close()
		defer conn.Close()
		io.Copy(client, conn)
	}()
	go func() {
		defer client.Close()
		defer conn.Close()
		io.Copy(conn, client)
	}()
}

func main() {

	listenPort := flag.String("listenPort", "80", "listenPort")
	clientPort := flag.String("clientPort", "8080", "clientPort")

	flag.Parse()

	listener, err := net.Listen("tcp", ":" + *listenPort)
	if err != nil {
		log.Fatalf("Failed to setup listener: %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("ERROR: failed to accept listener: %v", err)
		}
		log.Printf("Accepted connection %v\n", conn)
		go forward(conn, ":" + *clientPort)
	}

}
