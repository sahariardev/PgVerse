package core

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"
)

func TestStartServer(t *testing.T) {
	port := "5435"
	go func() {
		err := StartServer(port, func(conn net.Conn) {
			buf := make([]byte, 1024)

			for {
				n, err := conn.Read(buf)

				if err != nil {
					if err.Error() != "EOF" {
						fmt.Printf("Error reading from client: %v\n", err)
					}
					break // Exit the loop if there's an error or EOF
				}

				_, err = conn.Write(buf[:n])

				if err != nil {
					t.Errorf("Error writing to client: %v\n", err)
				}
			}
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
