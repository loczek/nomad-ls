package pipe

import (
	"io"

	"github.com/Ne0nd0g/npipe"
)

func GetTransport(pipe string) io.ReadWriteCloser {
	conn, err := npipe.Dial(pipe)
	if err != nil {
		panic(err)
	}

	return conn
}
