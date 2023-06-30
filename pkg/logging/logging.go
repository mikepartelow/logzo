package logging

import (
	"os"

	"github.com/mikepartelow/logzo/pkg/multipass"
	"github.com/mikepartelow/logzo/pkg/tater"
	"golang.org/x/exp/slog"
)

var tater_ *tater.Tater = nil

func Init(filename string) {
	if tater_ != nil {
		panic("re-init")
	}

	level := slog.LevelDebug

	var h slog.Handler = slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: level})

	if filename != "" {
		tater_ = tater.New(filename)

		h = multipass.New(h, slog.NewJSONHandler(tater_, &slog.HandlerOptions{Level: level}))
	}

	slog.SetDefault(slog.New(h))
}

func Rotate() error {
	return tater_.Rotate()
}
