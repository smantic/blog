---
title: "dependency injection" 
date: 2022-06-21T02:08:27-05:00
draft: false
tags: 
  - go
---

**In go** Dependency Injection is a slightly different way to pass arguments to your function.

* It's not a 'hack' or dangerous.
* It's not cargo cult programming.
* It's not complex.
* It can be done without a framework.
* You don't have to use it. But it will make your life easier. 

---

## Injection

The pinnacle example for dependency injection is when writing http handler functions using `net/http`. 

In go a function that handles http requests looks like this: 

```go 
func ServeHTTP(w http.ResponseWriter, r *http.Request){
    return 
}
```

This function signature is required in order to mount it onto the stdlib's http server.[^1] 

```go 
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func main() {

	s := http.NewServeMux()
	s.Handle("/foo/bar", http.HandlerFunc(HelloHandler))

	err := http.ListenAndServe(":8000", s)
	if err != nil {
		log.Println(err)
	}
}

```

We mounted our handler function ok, 
but what if I want to pass some more arguments to the function? 

For example, lets say I want to have a cli argument that changes some behavior in the handler. 
We can't exactly modify the signature of the function to add extra arguments, 
but we can capture extra arguments through dependency injection, without actually changing the function signature. 

```go 
type HelloHandler struct {
	data string
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, h.data)
}

func main() {

	var data string = "Hello, World!\n"

	if len(os.Args) == 2 {
		data = os.Args[1]
	}

	h := &HelloHandler{
		data: data,
	}

	s := http.NewServeMux()
	s.Handle("/foo/bar", h)

	err := http.ListenAndServe(":8000", s)
	if err != nil {
		log.Println(err)
	}
}
```

We've essentially added more arguments to our http handler function, without actually changing the signature of the function. 

Yes, You could just read os.Args[1] from the handler, but doing it this way tells the reader that this function requires a string to operate (in the same way a function argument does), 
It doesn't matter where you get it from; be it a file, the environment, or otherwise. As long as it's passed in as a string. 

This idea of 'declaring your dependencies` is also useful in testing. 


## Testing // Removing global state

Imagine a http handler writen like the following. Notice the use of an external package's global state. 

```go 

func FooHandler(w http.ResponseWriter, req *http.Request) {
    err := external.DefaultService.DoSomething()
    if err != nil { 
        http.Error(w, err, 500)
    }

    w.WriterHeader(200)
}
```

Unfortuantely, this external function call makes a http call to an unreliable 3rd party service, and somtimes fails. 
Which means if I were to write a test for this function it will fail randomly. 
We'd prefer to avoid that, so lets use dependency injection (and a interface) so we can replace this external call with a mock in our uint tests.

```go 
type DoSomethinger interface{ 
    DoSomething() error
}

type FooHandler struct { 
    dep DoSomethinger
}

func FooHandler(w http.ResponseWriter, req *http.Request) {
    err := external.DefaultService.DoSomething()
    if err != nil { 
        http.Error(w, err, 500)
    }

    w.WriterHeader(200)
}
```

So we would have a test that looks like this. 

```go 
func TestFooHandler(t *testing.T) { 
     m := NewMockDoSomethinger()
     // expect a call to DoSomething()

     f := FooHandler{ dep: m } 
     s := httptest.NewServer(http.HandlerFunc(f.FooHandler))

     r := httptest.NewRequest(http.MethodGET, s.URL, nil)
     resp, err := http.DefaultClient.Do(r)
     if err != nil { 
        t.Fatal(err)
     }

     if resp.StatusCode != 200 { 
        t.Error(err)
     }
}

```


<!---
## Testing 


Dependency Injection is also very useful in testing when we want to sub in mocks for out unit tests.  


Imagine we have a "Serivce" struct defined as follows. 
```go 
package main

import "io"

// Foo is something that we can write to.
type Foo interface {
	io.Writer
}

type Service struct {
	f Foo
}

func (s *Service) Bar() error {
	_, err := s.f.Write(nil)

    if err == io.EOF { 
        return nil
    }
	return err
}
```

It has a single dependency of an something that implements `io.Writer`.  
  
Now, we want to test `Bar()`. 
We want to make sure that we dont return an error if there isn't one from our dependency, and that we dont return io.EOF; 
Which we'll say is an implementation detail that the caller of `Bar()` doesn't care about. 

To create that scenario we need a mock `io.Writer` that will help us create those specific conditions. 

```go 
func TestBar(t *testing.T) {

	var mockFoo = io.Discard

	s := Service{
		foo: mockFoo,
	}

	got := s.Bar()
	if got != nil {
		t.Fatal(got)
	}
}
```
-->

[^1]: https://pkg.go.dev/net/http#Handler


