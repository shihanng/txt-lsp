package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cockroachdb/errors"
	lsp "github.com/sourcegraph/go-lsp"
	"github.com/sourcegraph/jsonrpc2"
)

func handle(_ context.Context, _ *jsonrpc2.Conn, req *jsonrpc2.Request) (interface{}, error) {
	switch req.Method {
	case "initialize":
		return lsp.InitializeResult{
			Capabilities: lsp.ServerCapabilities{
				CompletionProvider: &lsp.CompletionOptions{},
			},
		}, nil
	case "textDocument/completion":
		return lsp.CompletionList{
			IsIncomplete: false,
			Items: []lsp.CompletionItem{
				{Label: "Johor Johor (Johor Bahru)"},
				{Label: "Kedah Kedah (Alor Setar)"},
				{Label: "Kelantan Kelantan (Kota Bharu)"},
				{Label: "Malacca Malacca (Malacca City)"},
				{Label: "Negeri Sembilan Negeri Sembilan (Seremban)"},
				{Label: "Pahang Pahang (Kuantan)"},
				{Label: "Penang Penang (George Town)"},
				{Label: "Perak Perak (Ipoh)"},
				{Label: "Perlis Perlis (Kangar)"},
				{Label: "Selangor Selangor (Shah Alam)"},
				{Label: "Sabah Sabah (Kota Kinabalu)"},
				{Label: "Sarawak Sarawak (Kuching)"},
			},
		}, nil
	case "initialized":
		return nil, nil //nolint:nilnil
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
