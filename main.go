package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
	"runtime/debug"
	"strings"

	"github.com/lmittmann/tint"
	"github.com/loczek/nomad-ls/internal/lsp"
	"github.com/loczek/nomad-ls/internal/pipe"
	"go.lsp.dev/jsonrpc2"
)

type Flags struct {
	logLevel string
	stdio    bool   // stdin/stdout
	pipe     string // named pipe (Windows) or unix socket (Linux, Mac)
	socket   string // tcp socket port
}

var flags = Flags{}

func init() {
	flag.StringVar(&flags.logLevel, "log-level", "info", "language server log level")
	flag.BoolVar(&flags.stdio, "stdio", false, "stdin/stdout as the transport method")
	flag.StringVar(&flags.pipe, "pipe", "", "named pipe (Windows) or unix socket (Linux, Mac) as the transport method")
	flag.StringVar(&flags.socket, "socket", "", "port of the tcp socket as the transport method")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: nomad-ls [options]\n\n")
		fmt.Fprintf(os.Stderr, "Note: \"--stdio\", \"--pipe=...\" or \"--port=...\" must be defined\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()
}

func main() {
	w := os.Stderr

	var handler slog.Handler
	handlerOpts := slog.HandlerOptions{
		Level: slog.LevelError,
	}

	switch flags.logLevel {
	case "debug":
		handlerOpts.Level = slog.LevelDebug
	case "info":
		handlerOpts.Level = slog.LevelInfo
	case "warn":
		handlerOpts.Level = slog.LevelWarn
	case "error":
		handlerOpts.Level = slog.LevelError
	default:
		panic("invalid log level")
	}

	if isBuilt() {
		handler = slog.NewTextHandler(w, &handlerOpts)
	} else {
		handler = tint.NewHandler(w, &tint.Options{
			AddSource:   handlerOpts.AddSource,
			Level:       handlerOpts.Level,
			ReplaceAttr: handlerOpts.ReplaceAttr,
		})
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)

	var transport *rwc

	if isFlagPassed("pipe") {
		slog.Info("transport pipe", slog.String("pipe", flags.pipe))
		conn := pipe.GetTransport(flags.pipe)
		transport = &rwc{conn, conn}
	} else if isFlagPassed("socket") {
		slog.Info("transport socket", slog.String("socket", fmt.Sprintf(":%s", flags.socket)))
		conn, err := net.Dial("tcp", fmt.Sprintf(":%s", flags.socket))
		if err != nil {
			panic(err)
		}
		transport = &rwc{conn, conn}
	} else {
		slog.Info("transport stdio")
		transport = &rwc{os.Stdin, os.Stdout}
	}

	stream := jsonrpc2.NewStream(transport)
	con := jsonrpc2.NewConn(stream)

	service := lsp.New(con, *logger)

	con.Go(context.Background(), func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
		go func() {
			logger.Info("Received request", slog.String("method", req.Method()))

			resp, err := service.Handle(ctx, reply, req)

			logger.Info("response", "data", resp)

			reply(ctx, resp, err)

			if err != nil {
				logger.Error("Received error from handler", "error", err.Error())
			}
		}()
		return nil
	})

	logger.Info("Started", "build", buildInfo())

	<-con.Done()

	logger.Info("Exited")
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

func buildInfo() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}

	return info.Main.Version
}

func isFlagPassed(name string) bool {
	found := false

	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})

	return found
}
