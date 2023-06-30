package multipass

import (
	"context"

	"golang.org/x/exp/slog"
)

// MultiPass implements slog.Handler
type MultiPass struct {
	handlers []slog.Handler
}

func New(handlers ...slog.Handler) *MultiPass {
	hh := make([]slog.Handler, len(handlers))

	for i, h := range handlers {
		hh[i] = h
	}

	return &MultiPass{
		handlers: hh,
	}
}

func (mp *MultiPass) Enabled(context.Context, slog.Level) bool {
	return true
}

func (mp *MultiPass) Handle(ctx context.Context, r slog.Record) error {
	var errTop error
	for _, h := range mp.handlers {
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
