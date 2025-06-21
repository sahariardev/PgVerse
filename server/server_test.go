package server

import (
	"bufio"
	"net"
	"strings"
	"testing"
	"time"
)

func TestStartServer(t *testing.T) {
	port := "5433"
	go func() {
		err := StartServer(port, func(bytes []byte) []byte {
			return bytes
		})

		if err != nil {
			t.Errorf("Server error: %v", err)
		}
	}()

	time.Sleep(200 * time.Millisecond)

	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}

	defer conn.Close()

	requestBody := "Hello World!"

	_, err = conn.Write([]byte(requestBody))
	if err != nil {
		t.Fatalf("Failed to write to server: %v", err)
	}

	reader := bufio.NewReader(conn)
	resp, err := reader.ReadString('!')
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	if strings.TrimSpace(resp) != requestBody {
		t.Errorf("Expected %q, got %q", requestBody, resp)
	}
}
