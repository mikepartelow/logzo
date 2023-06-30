package logging

import (
	"io"
	"os"

	"github.com/mikepartelow/logzo/pkg/tater"
	"golang.org/x/exp/slog"
)

var t *tater.Tater = nil

func Init(filename string) {
	if t != nil {
		panic("re-init")
	}

	var out io.Writer = os.Stderr

	if filename != "" {
		t = tater.New(filename)

		out = io.MultiWriter(os.Stderr, t)
	}

	h := slog.NewJSONHandler(out, &slog.HandlerOptions{Level: slog.LevelDebug})
	slog.SetDefault(slog.New(h))
}

func Rotate() error {
	return t.Rotate()
}
