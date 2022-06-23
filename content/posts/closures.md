---
title: "injection with closures"
date: 2021-11-27T13:52:39-06:00
draft: false
tags: 
  - go
---

In languages that have first class functions we can define closures that lexically capture values in the scope that they are defined in.
  
```go 
package main

import "fmt"

func main() {
	b := true
	closure := func() error {
		err := fmt.Errorf("something bad happened!!")
		if b {
			return err
		}
		return nil
	}
	b = false

	err := closure()
	fmt.Printf("%v\n", err)
}
```
In this example we lexically capture `b` in our closure. Our function returns an error 
 if b is true. The value of b changes to false, causing the function to not return an error when invoked.

Lexically capturing variables can be useful in cases where you want to inject a variable into some function that doesn't accept that variable as a part of it's signature. 
Such as a function that accepts only `net/http.Handler` [^1] as an input.

Here is an example where create a closure, lexically capturing a `uber/zap.Logger`, and using it to log a request after it's finished.

```go 
type ZapMiddleware struct {
	logger *zap.Logger
}

func (z *ZapMiddleware) LogRequest(h http.Handler) http.Handler {

	f := func(w http.ResponseWriter, r *http.Request) {

		wr := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		start := time.Now()

		h.ServeHTTP(wr, r)

		dur := time.Since(start)
		z.logger.Info(
			"request-finished",
			zap.Int("status", wr.Status()),
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.Duration("duration", dur),
			zap.String("user-agent", r.UserAgent()),
		)
	}
	return http.HandlerFunc(f)
}
```

Then the function gets passed to the router to be used as middleware.

We can also use this trick to inject `context.Context` into functions that don't accept context as part of their signature.

The signature for a `bwmarrin/discordgo` interaction handler is `func(s *discordgo.Session, i *discordgo.InteractionCreate)` [^2]
We can't directly pass in the context, but we can capture it as a part of the closure.

```go

func (d *Discord) HandleInteraction(ctx context.Context)  func(s *discordgo.Session, i *discordgo.InteractionCreate) {

	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
        err := d.Add(ctx, s, i)
        if err != nil {
            log.Println(err)
        }
    }
}

// Init starts the discord service and adds handlers.
func (d *Discord) Init(ctx context.Context) error {

	d.session.AddHandler(d.HandleInteraction(ctx))

    ...
}
```

In this example I am passing context to a function that returns back a closure to be used as a handler.



[^1]: https://pkg.go.dev/net/http#Handler
[^2]: https://github.com/bwmarrin/discordgo/blob/c27ad65527ecbc264c674cd3d0e85bb09de942e3/eventhandlers.go#L916
