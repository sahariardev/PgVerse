package core

import (
	"github.com/jackc/pgproto3/v2"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	frontend := pgproto3.NewFrontend(pgproto3.NewChunkReader(conn), conn)
	reader := pgproto3.NewFrontend(pgproto3.NewChunkReader(conn), conn)

}
