package main

import (
	"io"
	"log"
	"net"
	"time"
)

// Define a type for the connection handler function
type connectionHandler func(net.Conn, string)

func handleConnection(clientConn net.Conn, targetAdd string) {
	defer clientConn.Close()

	targetConn, err := net.Dial("tcp", targetAdd)
	if err != nil {
		log.Println("Failed to connect to target server", targetAdd)
	}
	defer targetConn.Close()

	go io.Copy(targetConn, clientConn)
	io.Copy(clientConn, targetConn)
	time.Sleep(4 * time.Second) // Sleep for 4 seconds

}

func listenAndHandleConnection(listenAddr string, targetAddr string, handler connectionHandler) {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v \n", listenAddr, err)
	}
	defer listener.Close()

	log.Println("Proxy is now listening on", listener.Addr().String())

	maxConnections := 3
	sem := make(chan struct{}, maxConnections)

	for {
		// Trying to acquire the semaphore and handle the connection
		sem <- struct{}{}

		clientConn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection", err)
			continue
		}

		log.Println("Connection accepted from", clientConn.RemoteAddr().String())

		go func() {
			handler(clientConn, targetAddr)
			// Release the semaphore slot
			<-sem // Release the semaphore slot before finishing the handler
		}()
	}
}

func main() {
	listenAddr := "localhost:8080"
	targetAddr := "httpbin.org:80"

	listenAndHandleConnection(listenAddr, targetAddr, handleConnection)

}
