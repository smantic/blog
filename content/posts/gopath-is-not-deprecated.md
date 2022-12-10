---
title: "gopath is not deprecated"
date: 2022-12-09T21:08:24-06:00
draft: false
tags: 
  - go
---

Maybe you've heard someone say something like; 
> "Gopath is deprecated, use go modules".

While what they mean to say is true, developing in gopath mode is deprecated. However, when interpreted literally this statement is not true. And very likely confuses new comers to go. 
Lets peel back a layer of the onion and be exact with what we mean to say. 

The action of developing inside of Gopath *and* not using modules is deprecated. 
This is (uncommonly) called 'gopath development mode' or just 'gopath mode' and is how development was done in go prior to go modules. 
That has all changed when go got modules (go's dependency management system). 
You can still write code utilizing gopath mode, however it is discouraged as at this point you should be using go modules in all of your go projects. 
Thus why we say that developing in gopath mode is deprecated. 

Maybe this confusion stirs from a few poorly worded google results. [1](https://github.com/golang/go/issues/30329), [2](https://pagure.io/GoSIG/go-sig/issue/35), [3](https://groups.google.com/g/Golang-dev/c/hGwvCceDr140)



So gopath *mode* is deprecated. But what is gopath, and what is it used for? 

Gopath is an environment variable that which defaults to $HOME/go. 
And is used to determine defaults for where modules are downloaded to($GOMODCACHE), where binaries are installed ($GOBIN), and where checksums are stored ($GOPATH/pkg/sumdb).

So gopath is itself is still very much in use, and will continue to be in the future. 

All this info and more can be found in this very helpful [wiki](https://github.com/golang/go/wiki/GOPATH))! 

Read More 
----
* [Gopath wiki](https://github.com/golang/go/wiki/GOPATH)
* [Go Modules Blog](https://go.dev/blog/modules2019)
