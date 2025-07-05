package core

import (
	"fmt"
	"io"
	"net"
	"sync"
)

func RequestHandler(clientConn net.Conn) {
	defer clientConn.Close()

	pgServerConn, err := net.Dial("tcp", "localhost:5432")

	if err != nil {
		fmt.Printf("Failed to connect to server: %v", err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		defer pgServerConn.Close()
		io.Copy(pgServerConn, clientConn)
	}()

	go func() {
		defer wg.Done()
		defer clientConn.Close()
		io.Copy(clientConn, pgServerConn)
	}()

	wg.Wait()
}
