---
title: "struct-to-map"
date: 2021-10-26T21:02:11-05:00
draft: false
toc: false
images:
tags: 
  - go  
  - benchmark
---

Occasionally in go we want convert our structured data into an unstructured format, such as a `map[string]interface{}`

for struct like: 

```go 
type Foo struct {
	A Bar    `json:"a"`
	B Bar    `json:"b"`
	C Bar    `json:"c"`
	D Bar    `json:"d"`
	E Bar    `json:"e"`
	F int    `json:"f"`
	G uint   `json:"g"`
	H string `json:"h"`
}

type Bar struct {
	Bar string `json:"bar"`
}
```

There are a couple ways we can convert it into a map. 

1. manually 

```go 
func base(foo Foo) map[string]interface{} {

	result := make(map[string]interface{})

	result["A"] = foo.A
	result["B"] = foo.B
	result["C"] = foo.C
	result["D"] = foo.D
	result["E"] = foo.E
	result["F"] = foo.F
	result["G"] = foo.G
	result["H"] = foo.H
	return result
}
```

just manually assigning the fields and keys.

2. reflection 

We can use reflection to generically get the name of the struct fields and their data.

```go 
func reflec(foo Foo) map[string]interface{} {

	result := make(map[string]interface{})

	val := reflect.ValueOf(foo)
	typ := reflect.TypeOf(foo)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		result[typ.Field(i).Name] = field.Interface()
	}

	return result
}
```

3. `encoding/json` hack

A more established go dev may be tempted to try the following:

```go 
func hack(foo Foo) map[string]interface{} {

	result := make(map[string]interface{})

	bytes, _ := json.Marshal(foo)
	json.Unmarshal(bytes, result)

	return result
}
```

Much shorter than option 1, but we'll find it's much less efficient.  


---

## Results

We can benchmark these functions: 

```go
// to elim some compiler optimizations
var result map[string]interface{}

func BenchmarkBase(b *testing.B) {
	var m map[string]interface{}
	for n := 0; n < b.N; n++ {
		m = base(foo)
	}
	result = m
}
```

Benchmarking on my machine I get results like: 
```
$ go test -bench=. -benchtime=60s

goos: linux
goarch: amd64
pkg: github.com/smantic/bench-struct-to-map
cpu: Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz
BenchmarkBase-8      	133096639	       538.1 ns/op
BenchmarkHack-8      	47031236	      1591 ns/op
BenchmarkReflect-8   	72736426	       965.5 ns/op
PASS
ok  	github.com/smantic/bench-struct-to-map	273.862s
```

We see that the option 1 has a speedup of ~2x compared to the reflect method, and ~3x speedup when compared to the encoding/json hack. 

Option 1 is probably faster due to the following ideas: 
 1. runtime reflection is slow   
 2. encoding/json has extra logic for json handling
    * we are also doing an extra step to convert the struct to a byte array of json data.  
    * encoding/json uses reflection to do this in the first place.  

The base option is just assigning entries in the map, and lets the compiler handle the rest.

Makes me feel a bit better about writing out the struct mapping by hand!

code : <https://github.com/smantic/bench-struct-to-map>


--- 
### Other Libraries 

If you google "go struct to map" you are likely to stumble upon [this SO post](https://stackoverflow.com/questions/23589564/function-for-converting-a-struct-to-map-in-golang).   
The accepted answer mentions [fatih/structs](https://github.com/fatih/structs). However, this project still uses reflection and has been unmaintained for years. 

Any other project will likely have performance worse or close to the reflection method, since they will all be using reflection underneath. 

You may also find [mitchellh/mapstructure](https://github.com/mitchellh/mapstructure), which still uses reflection.   
For fun here's the benchmark for that.

``` 
$ go test -bench=Mitchellh -benchtime=60s

goos: linux
goarch: amd64
pkg: github.com/smantic/bench-struct-to-map
cpu: Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz
BenchmarkMitchellh-8   	11661878	      6298 ns/op
PASS
ok  	github.com/smantic/bench-struct-to-map	79.679s
```
  ¯\\\_(ツ)\_/¯

