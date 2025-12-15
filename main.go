package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime/debug"
	"strings"

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

	service := lsp.New(con, *logger)

	con.Go(context.Background(), func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
		go func() {
			logger.Info(fmt.Sprintf("recieved method: %s", req.Method()))

			resp, err := service.Handle(ctx, reply, req)

			logger.Info("response", "data", resp)

			reply(ctx, resp, err)

			if err != nil {
				logger.Info("received error from handler", "error", err.Error())
			}
		}()
		return nil
	})

	logger.Info("starting", "build", BuildInfo())

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

func BuildInfo() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}
	var rev string
	for _, s := range info.Settings {
		if s.Key == "vcs.revision" {
			rev = s.Value[:min(7, len(rev))]
			break
		}
	}
	return fmt.Sprintf("%s-%s", info.Main.Version, rev)
}
