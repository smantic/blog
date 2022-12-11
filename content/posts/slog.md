---
title: "x/slog"
date: 2022-12-09T21:57:50-06:00
draft: false
tags: 
  - go
---

`golang.org/x/exp/slog` is a new experimental package implementing a structured logger. Originating from this [proposal](https://go.googlesource.com/proposal/+/master/design/56345-structured-logging.md) it aims to get a structured logging library into the stdlib. 

If you're not familiar with structured logging. It is, generally, the superior way to do logging. The idea is just to add some structure to the log lines in the form of key value pairs. 
Typical structure is usually either [JSON format](https://pkg.go.dev/golang.org/x/exp/slog#JSONHandler), or simply [separated by a space](https://pkg.go.dev/golang.org/x/exp/slog#TextHandler).


I've recently been working on a logging middleware to use at work that utilizes x/slog. It looks like this.

```go
	l := mw.slog.With(
		slog.String("path", path),
		slog.String("route", route),
		slog.Time("start", start),
	).WithContext(ctx)

	ww := &middleware.ResponseWriter{ResponseWriter: w}
	r = r.WithContext(slog.NewContext(ctx, l))

	next.ServeHTTP(ww, r)

	headers := mw.getHeaderAttributes(r)
	queries := mw.getURLQueryAttributes(r)

	finish := time.Now()
	duration := time.Since(start)

	l.Info(
		"finished http request",
		slog.Duration("duration", duration),
		slog.Time("finish", finish),
		slog.Int("status", ww.StatusCode),
		slog.Int("bytes-written", ww.BytesWritten),
		slog.Group("headers", headers...),
		slog.Group("url-queries", queries...),
	)
```

The first part creates a new slog logger with some fields to start with.
It creates a new logger from our base logger stored in the middleware struct `mw` (see: [dependency injection](https://blog.smantic.dev/posts/dependency-injection/)). We also attach our request context to the logger. This can be helpful in the case your logging backend stores fields in fields in the context, so you can pass the request context to the call to your logging backend.  

Next I create a new context with [slog.NewContext](https://pkg.go.dev/golang.org/x/exp/slog#NewContext), which gives you a context that contains your logger. 
This is so that downstream handlers from this logging middleware can access the logger (and any fields you've set on the logger). They can get it by using [slog.FromContext](https://pkg.go.dev/golang.org/x/exp/slog#NewContext). This is context based dependency injection and is common strategy when dealing with meta or life cycle things like logging and is particularly useful to propagate these things between middleware and http handlers.   

To give you an idea of what this is like, here are 2 example log lines from this middleware. The initial fields path, route, and started got attached to the logger, so all downstream log calls includes those fields.   
```
2022/12/10 18:15:35 INFO log line for event that happens durring the handling of the request path=/test route=test start=2022-12-10T18:15:35.774-06:00
2022/12/10 18:15:35 INFO finished http request path=/test route=test start=2022-12-10T18:15:35.774-06:00 duration=120.542µs finish=2022-12-10T18:15:35.774-06:00
```

However all of this so far are typical of any structured logger that we already have. 
So lets talk about some of the things that makes x/slog different.

slog's interface to create a log line is as follows. 
```go 
func (l *Logger) Log(level Level, msg string, args ...any)
```

Its pretty common for a logger to accept a vardic of fields to log, and the idea is that you are passing in [slog.Attr](https://pkg.go.dev/golang.org/x/exp/slog#Attr) 
to be logged as key value pairs. But you don't have to. In fact, as the signature of `Log` suggests, you can pass in any type of thing into the function call. 
```go
	l.Info(
		"finished http request",
		"key",
		"value",
	)
```
produces this 
```
2022/12/10 18:27:02 INFO finished http request key=value
```

The call will log will pair up key value fields given you pass an even number of non attribute types into the function. 
You might be wondering, what if you dont? 
```go 
2022/12/10 18:30:08 INFO finished http request !BADKEY=key
```

This seems like a potentially annoying trade off (untyped convenience vs potential bad number of arguments). 
I'm a pretty a big fan of [uber/zap's](https://pkg.go.dev/go.uber.org/zap#Logger.Log) log interface, maybe they will adopt something similar? 

---- 

There is one thing that I really like about x/slog is the ability to change the logging backend that it uses. This means that slog.Logger can be a unifying logging interface (for structured logging). Something that go has struggled with for a long time. 

This means I can write a library. Say a logging middleware. The logging middleware can accept a
`*slog.Logger` as a dependency. And you can pass in whatever logger you want into slog's [New function](https://pkg.go.dev/golang.org/x/exp/slog#New), and pass me the resulting `*slog.Logger`. 


This will be useful for me at work, as I can write a wrapper around our company's logger. But here is an example that creates a handler wrapper around uber/zap. 

```go
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
    //TODO
	return nil
}
func (zs *ZapSlogWrapper) WithGroup(n string) slog.Handler {
    // TODO
	return nil
}

func main() {

	z, _ := zap.NewProduction()
	zw := ZapSlogWrapper{log: z}
	s := slog.New(&zw)
	mw := Middleware{s}
	mw.slog.Info("test")
}

```

```
{"level":"info","ts":1670721201.439241,"caller":"slog/main.go:24","msg":"test"}
```


