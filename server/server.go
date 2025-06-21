package server

import (
	"fmt"
	"net"
)

type ResponseHandler func([]byte) []byte

func StartServer(port string, handler ResponseHandler) error {
	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		panic(err)
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			panic(err)
		}
	}(listener)

	for {
		conn, err := listener.Accept()

		if err != nil {
			return err
		}

		go func(conn net.Conn) {
			defer func(conn net.Conn) {
				_ = conn.Close()
			}(conn)

			buff := make([]byte, 1024)

			for {
				n, err := conn.Read(buff)
				if err != nil {
					return
				}

				fmt.Print(string(buff[:n]))

				response := handler(buff[:n])

				if response != nil {
					_, err = conn.Write(handler(buff[:n]))

					if err != nil {
						return
					}
				}
			}
		}(conn)

	}

	return nil
}
