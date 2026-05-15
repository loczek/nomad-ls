package pipe

import (
	"io"
	"net"
)

func GetTransport(pipe string) io.ReadWriteCloser {
	conn, err := net.Dial("unix", pipe)
	if err != nil {
		panic(err)
	}

	return conn
}
