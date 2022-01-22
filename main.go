package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cockroachdb/errors"
	lsp "github.com/sourcegraph/go-lsp"
	"github.com/sourcegraph/jsonrpc2"
)

func handle(_ context.Context, _ *jsonrpc2.Conn, req *jsonrpc2.Request) (result interface{}, err error) {
	switch req.Method { //nolint:gocritic
	case "initialize":
		return lsp.InitializeResult{}, nil
	}

	return nil, &jsonrpc2.Error{
		Code:    jsonrpc2.CodeMethodNotFound,
		Message: fmt.Sprintf("method not supported: %s", req.Method),
	}
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
