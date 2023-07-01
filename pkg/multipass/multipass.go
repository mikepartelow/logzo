package multipass

import (
	"context"

	"golang.org/x/exp/slog"
)

// MultiPass implements slog.Handler
type MultiPass struct {
	handlers []slog.Handler
	// FIXME: buffered
	channels []chan slog.Record
}

func New(handlers ...slog.Handler) *MultiPass {
	hh := make([]slog.Handler, len(handlers))
	cc := make([]chan slog.Record, len(handlers))

	for i, h := range handlers {
		hh[i] = h
		cc[i] = make(chan slog.Record)
		go func(handler slog.Handler, channel chan slog.Record) {
			for r := range channel {
				// FIXME: handle error
				// FIXME: receive context from MultiPass.Handle
				_ = handler.Handle(context.Background(), r)
			}
		}(hh[i], cc[i])
	}

	return &MultiPass{
		handlers: hh,
		channels: cc,
	}
}

func (mp *MultiPass) Enabled(context.Context, slog.Level) bool {
	return true
}

func (mp *MultiPass) Handle(ctx context.Context, r slog.Record) error {
	for _, c := range mp.channels {
		c <- r
	}
	// FIXME: receive errors from goroutines
	return nil
}

func (mp *MultiPass) WithAttrs(attrs []slog.Attr) slog.Handler {
	panic("NIH")
}
func (mp *MultiPass) WithGroup(name string) slog.Handler {
	panic("NIH")
}
