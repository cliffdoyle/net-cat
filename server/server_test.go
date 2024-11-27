package server

import (
	"net"
	"testing"
	"time"
)

func TestStartServer(t *testing.T) {
	// Temporarily run the server in a goroutine to avoid blocking the test
	go func() {
		StartServer("8989")
	}()

	// Give the server a little time to start up
	time.Sleep(1 * time.Second)

	// Now check if we can connect to the server
	conn, err := net.Dial("tcp", "localhost:8989")
	if err != nil {
		t.Fatalf("Error connecting to server: %v", err)
	}
	defer conn.Close()

	// Send a message to the server
	conn.Write([]byte("Hello Server!"))
}
