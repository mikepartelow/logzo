package logging

import (
	"context"
	"os"

	"github.com/mikepartelow/logzo/pkg/tater"
	"golang.org/x/exp/slog"
)

var tater_ *tater.Tater = nil

// MultiPass implements slog.Handler
type MultiPass struct {
	Handlers []slog.Handler
}

func (mp *MultiPass) Enabled(context.Context, slog.Level) bool {
	return true
}

func (mp *MultiPass) Handle(ctx context.Context, r slog.Record) error {
	var errTop error
	for _, h := range mp.Handlers {
		// FIXME: goroutines
		if err := h.Handle(ctx, r); err != nil {
			errTop = err
		}
	}
	return errTop
}

func (mp *MultiPass) WithAttrs(attrs []slog.Attr) slog.Handler {
	panic("NIH")
}
func (mp *MultiPass) WithGroup(name string) slog.Handler {
	panic("NIH")
}

func Init(filename string) {
	if tater_ != nil {
		panic("re-init")
	}

	level := slog.LevelDebug

	mp := &MultiPass{
		Handlers: []slog.Handler{
			slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: level}),
		},
	}

	if filename != "" {
		tater_ = tater.New(filename)
		mp.Handlers = append(mp.Handlers,
			slog.NewJSONHandler(tater_, &slog.HandlerOptions{Level: level}),
		)
	}

	slog.SetDefault(slog.New(mp))
}

func Rotate() error {
	return tater_.Rotate()
}
