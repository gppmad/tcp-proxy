package main

import (
	"io"
	"log"
	"net"
)

func handleConnection(clientConn net.Conn, targetAdd string) {
	defer clientConn.Close()

	targetConn, err := net.Dial("tcp", targetAdd)
	if err != nil {
		log.Printf("Failed to connect to target server", targetAdd)
	}
	defer targetConn.Close()

	go io.Copy(targetConn, clientConn)
	io.Copy(clientConn, targetConn)

}

func main() {
	listenAddr := "localhost:8080"
	targetAddr := "httpbin.org:80"

	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal("Failed to listen on %s \n", listenAddr, err)
	}
	defer listener.Close()

	log.Println("Proxy is now listening on", listenAddr)
	for {
		clientConn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection", err)
			continue
		}
		go handleConnection(clientConn, targetAddr)
	}

}
