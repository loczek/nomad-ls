package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/lmittmann/tint"
	"github.com/loczek/nomad-ls/internal/lsp"
	"go.lsp.dev/jsonrpc2"
)

func main() {
	w := os.Stderr

	handler := tint.NewHandler(w, nil)
	if isBuilt() {
		handler = slog.NewTextHandler(w, nil)
	}

	logger := slog.New(handler)

	stream := jsonrpc2.NewStream(&rwc{os.Stdin, os.Stdout})
	con := jsonrpc2.NewConn(stream)

	lsp := lsp.New(con, *logger)

	parser := hclparse.NewParser()

	file, err := os.ReadFile("./loki.nomad.hcl")
	if err != nil {
		panic(err)
	}

	parser.ParseHCL(file, "loki")

	body := parser.Files()["loki"].Body
	// file_bytes := parser.Files()["loki"].Bytes
	// logger.Info(fmt.Sprintf(string(file_bytes)))

	logger.Info(fmt.Sprintf("body before: %#v", body))
	content, body, err := body.PartialContent(&hcl.BodySchema{})
	if err.Error() != "no diagnostics" {
		panic(err)
	}
	logger.Info(fmt.Sprintf("body after: %#v", body))
	logger.Info(fmt.Sprintf("body content: %#v", content))

	for _, block := range content.Blocks {
		logger.Info(fmt.Sprintf("block: %#v", block))
		logger.Info(fmt.Sprintf("block body: %#v", block.Body))
	}

	// x := decoder.Decoder{}
	// decoder.NewDecoder(lang.Path{
	// 	Path: "",
	// 	LanguageID: "nomad",
	// })

	con.Go(context.Background(), func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
		go func() {
			logger.Info(fmt.Sprintf("recieved method: %s", req.Method()))

			resp, err := lsp.Handle(ctx, reply, req)

			logger.Info("response: %#v", resp)

			reply(ctx, resp, err)

			if err != nil {
				logger.Info("recieved error from handler: %s", err)
			}
		}()
		return nil
	})

	logger.Info("starting")

	<-con.Done()

	logger.Info("exited")
}

type rwc struct {
	r io.ReadCloser
	w io.WriteCloser
}

func (rwc *rwc) Read(b []byte) (int, error)  { return rwc.r.Read(b) }
func (rwc *rwc) Write(b []byte) (int, error) { return rwc.w.Write(b) }
func (rwc *rwc) Close() error {
	rwc.r.Close()
	return rwc.w.Close()
}

func isBuilt() bool {
	entrypoint := string(os.Args[0])

	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}

	if strings.HasPrefix(entrypoint, os.TempDir()) || strings.HasPrefix(entrypoint, userCacheDir) {
		return false
	}

	return true
}
