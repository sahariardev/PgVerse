package core

import (
	"net"
)

type ResponseHandler func(conn net.Conn)

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
		go handler(conn)
	}

	return nil
}
