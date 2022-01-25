---
title: "errors"
date: 2022-01-24T18:59:48-06:00
draft: false
---

Runtime errors are a bit different in Go compared to other languages. Go lacks the familiar error handling like try/catch. 
Gophers think about errors a bit differently. In Go [errors are values](https://go.dev/blog/errors-are-values) like any other kind of type in go. 
Because of this we often see signatures that are like 

`func DoSomething() error`
or
`func FetchSomething() (results, error)`

And our functions tend to be composed of a lot of 
```go 
if err != nil { 
    return err
}
```
We call this "bubbling up" the error, as the error is bubbled up the call stack until we reach a point where we can make a decision about the error. 
Such as, logging the error and then exiting the program if it is a fatal error. 

The `error` type is defined as follows; 

```go 
type error interface {
    Error() string
}
```
Any type that implements the error interface can be considered an error.

Outside of implementing the error type the stdlib offers convenient ways to create error values through `errors.New(msg string)` and `fmt.Errorf(fmt string)`
The latter allows us to use formatting verbs in the error string, including the error wrapping format verb `%w`. It tends to be the more useful of the two.

```go
err := fmt.Errorf("failed to do %s because %w", something, otherErr)
```

The other frequently used package in go to wrap errors is through davecheney's  
`github.com/pkg/errors`
```go
errors.Wrap(err error, msg string)
``` 
which will wrap the error with a message *and a stack trace*. 
The wrapping of the error with a stack trace every time isn't all that useful, as Dave mentions in [this issue](https://github.com/pkg/errors/issues/245) looking for new maintainers.
Dave Cheney's package and perspective on error handling was once considered standard or idiomatic but the community (including him) has moved away from this practice. 
This package has been archived and you should stick with the stdlib's `errors` package.



## Panic 

Panics happen when the program crashes. If you index out of bounds, dereference a nil pointer or do other illegal operations your program will panic.
Panics escape the stack and immediately crash the program. You can invoke panic manually, 
but you should only do so if you actually want the program to crash, instead of gracefully handling the error.

## Recover

The only way to recover from a panic is to call `recover()` after the panic has happened. To be able to do this we have to push a recover onto the stack after any panic happens.
We can do this by using defer which allows us push a function call to the top of the stack. 

```go 
func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered from: ", r)
        }

        msg := r.(string)
    }()
    panic("something bad happened!")
}
```

The function pushes a function that calls recover to the top of the stack and then panics. 
The panic starts to unwind the existing call stack but the deferred function is yet to be executed. Recover gets called and captures the panic.
The value returned from recover is the value passed to panic, which is an `interface{}` value (or `any` if you're reading this post 1.18).

An astute developer may notice that it is possible to implement a sort of "try/catch" with panics, defer and recover, but you shouldn't do this.
The community is strongly in favor of thinking of errors as values as oppose to a try/catch strategy with runtime exceptions. 

----
### Further Reading
 * [Errors Are Values](https://go.dev/blog/errors-are-values) - Rob Pike
 * [Error Handling and Go](https://go.dev/blog/error-handling-and-go) - Andrew Gerrand
 * [Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover) - Andrew Gerrand
 * [On the uses and misuses of panics in Go](https://eli.thegreenplace.net/2018/on-the-uses-and-misuses-of-panics-in-go/) - Eli Bendersky
