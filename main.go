package main

import (
	"context"
	"os"

	"github.com/cockroachdb/errors"
	"github.com/sourcegraph/jsonrpc2"
)

var errNotImplemented = errors.New("not implemented")

func handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (result interface{}, err error) {
	return nil, errNotImplemented
}

func main() {
	handler := jsonrpc2.HandlerWithError(handle)

	connOpt := []jsonrpc2.ConnOpt{}

	<-jsonrpc2.NewConn(
		context.Background(),
		jsonrpc2.NewBufferedStream(stdrwc{}, jsonrpc2.VSCodeObjectCodec{}),
		handler, connOpt...).DisconnectNotify()
}

type stdrwc struct{}

func (stdrwc) Read(p []byte) (int, error) {
	n, err := os.Stdin.Read(p)

	return n, errors.Wrap(err, "main: read from stdin")
}

func (c stdrwc) Write(p []byte) (int, error) {
	n, err := os.Stdout.Write(p)

	return n, errors.Wrap(err, "main: write to stdout")
}

func (c stdrwc) Close() error {
	if err := os.Stdin.Close(); err != nil {
		return errors.Wrap(err, "main: close stdin")
	}

	return errors.Wrap(os.Stdout.Close(), "main: close stdout")
}
