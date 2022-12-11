package main

import (
	"net/http"
	"time"

	"go.uber.org/zap"
	"golang.org/x/exp/slog"
)

type Middleware struct {
	slog *slog.Logger
}

type ZapSlogWrapper struct {
	log *zap.Logger
}

func (zs *ZapSlogWrapper) Enabled(l slog.Level) bool {
	return true
}

func (zs *ZapSlogWrapper) Handle(r slog.Record) error {
	zs.log.Info(r.Message)
	return nil
}
func (zs *ZapSlogWrapper) WithAttrs(attrs []slog.Attr) slog.Handler {
	return nil
}
func (zs *ZapSlogWrapper) WithGroup(n string) slog.Handler {
	return nil
}

func main() {

	z, _ := zap.NewProduction()
	zw := ZapSlogWrapper{log: z}
	s := slog.New(&zw)
	mw := Middleware{s}
	mw.slog.Info("test")
}

func main2() {

	r, _ := http.NewRequest(http.MethodGet, "/test", nil)
	ctx := r.Context()

	start := time.Now()
	l := slog.With(
		slog.String("path", r.URL.Path),
		slog.String("route", "test"),
		slog.Time("start", start),
	).WithContext(ctx)

	r = r.WithContext(slog.NewContext(ctx, l))

	//next.ServeHTTP(ww, r)

	finish := time.Now()
	duration := time.Since(start)
	slog.Ctx(r.Context()).Info("log line for event that happens durring the handling of the request")

	_, _ = finish, duration
	f := ""
	l.Info(
		"finished http request",
		"key",
		struct{ foo *string }{&f},
	)
}
