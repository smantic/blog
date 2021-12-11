---
title: "Hitchhiker's Guide To Context"
date: 2021-12-09T18:13:26-06:00
draft: false
tags:
  - go
---

If you're new to Go you might have seen some functions like

```go 
func Foo(ctx context.Context) error
```

Context is a pattern frequently used in systems programming to help manage the life cycle of complex systems. 
We see often in Go as its considered good practice, is part of the standard library, and is extremely useful.

In go we have `context.Context` which is an interface like [^1]

```go 
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
``` 

However, we will mostly be working with the functions exposed by the package itself, namely:  

```go 
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key, val interface{}) Context
```

# Using contexts

Put simply. You should use context whenever a resource is being used, such as a database, file(s), a http server, and etc. 
Which happens to a very frequent thing that we do as programmers, and why we see it used all the time in go. 

### Cancellation

The bread and butter of `context.Context` is to be used to control the flow of the program by providing timeout and cancellation logic.  

In the case that something happens and we can to tell downstream consumers of the
context to halt, or cancel execution we can use the provided cancel function.

```go
import (
    "context" 
)

func SomeProcess(parentCtx context.Context) error { 
    ctx, cancel := context.WithCancel(parentCtx)
    
    go DoQuery(ctx)

    err := somethingElse()
    if err != nil { 
        cancel()
        return err
    }
    return nil
}
```

Something errors and we don't need to continue the work being done in `DoQuery`, so we signal the context to be cancelled by 
invoking the cancel function we got. When downstream consumers of the context inspect the context they will see that the context has been cancelled.

### Timeouts 

Timeouts are useful for situations where you may have a long running request or 
database query and you don't have all day to wait, or want to guard against an error that would cause 
the request to hang infinitely. 

We can create a context that has a timeout like:
```go 
import (
    "time"
    "context" 
)

func LongRuningDatabaseQuery(parentCtx context.Context) error { 

    /// create a new context that will timeout 1 minute from now.
    // or otherwise cancelled upstream
    ctx, cancel := context.WithTimeout(parentCtx, time.Minute)
    
    err := DoQuery(ctx)
    cancel()
    if err != nil { 
        return err
    }
    return nil
}
```

Note that `context.WithTimeout` gives us a cancel func like `context.WithCancel` does.
Its good practice to call cancel after the downstream consumer has finished, ensuring the newly created context does not leak.

### Consuming a context

Sometimes you may want to use context to let you know if you should continue with some operation in your system. 
How do we know when a context has been cancelled? Or its deadline has been passed? 

The way context implements _cancellation_ is by using a channel, when we receive something on the channel we know we should halt further work.
There are a couple ways you can do this. 

One way is by using `select`

```go
import( 
    "context"
)

func SomeRunningProcesses(ctx context.Context) error { 

    fooChan := someProcess()

    select { 
    case <-ctx.Done():
        return ctx.Error()
    case x <- fooChan: 
        return nil
    }
}
```

`select` will execute the case that happens first. Either we receive value `x` from the generated `fooChan`, 
or we receive a struct on the done channel of the context.
If the context has signaled done, then `context.Error()` will give us an error 
noting that the context was either cancelled or that the deadline has passed. 

If you are working in a single thread, and want to check that the context hasn't been 
cancelled yet periodically we can use `context.Error()` to check the state of the context.

```go 
for context.Error() == nil { 
    ...
}
```

### Logging 

With contexts we have the ability to attach values to the context using 
```go
func WithValue(parent Context, key, val interface{}) Context
```
We can add any key / value to the context to be passed to downstream consumers. 
Its not recommended to use this for critical types in your system, since we loose any strict typing, but
for non-critical things like logging, this feature is useful. 

```go
import "context"

func Foo(ctx context.Contex) { 
    
    var key string = "key"
    ctx = context.WithValue(ctx, key, "value")

    val := ctx.Value(key)
    if v, ok := val.(string); ok { 
        ...
    }
}
```

`WithValue` is also useful for keeping track of other non-critical state such as a trace/span IDs, or an event log.

### Traces and Spans

If your system is complex enough, it may be useful to implement traces and spans. 
Traces and spans give us valuable insights that are immediately apparent such as: what the call graph of our system looks like, and how long we are spending in each function.

If you use something such as lighshot you can get a very informative call graph of your trace.
{{< 
    figure src="/images/trace_main.webp"
    caption="A screenshot of lightstep's very nice trace overview"
    link="https://docs.lightstep.com/images/docs/trace_main.png?_cchid=a5af5c6653869d9aa940829e87bd5196"
>}}

### Parent contexts

So far in all of the examples we have a function that has a context provided to it. This is because the idea of context is that it starts at the beginning of your process.
When that is depends on what you are doing. 

If you are running a program that needs to cancel when an interrupt is received then you will start with `context.Background()` 
and attach the necessary signals to the context in the main thread.


```go 
import (    
    "context"
    "os/signal"
)

func main() { 
    // cancels ctx if os.Interrupt or os.Kill is recieved.
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
}
```
Sometimes you are working with a process spawned by a http server. 
In this case, we get the context attached to the http request since we are handling the request in a newly spawned thread.
We then pass that context to functions that the http handler calls.

```go
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request){ 
    ctx := r.Context()

    h.Service.SomeFunction(ctx)
}
```
### And More...

We've seen that contexts are a versatile pattern that enable a wide array of useful functionality, including Timeouts, Cancellations, Logging, and Tracing. 
I'm looking forward to seeing even more functionality being enabled by the use of context.

[^1]: https://pkg.go.dev/context#Context
