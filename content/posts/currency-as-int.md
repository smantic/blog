---
title: "ten-cents"
date: 2021-10-27T20:13:28-05:00
draft: true
toc: false
images:
tags: 
  - finance 
  - go
---

In programming we have floats and doubles to represent decimal values. 

It's very tempting to use these data types to represent money in software, since we often express money as whole units 
and a fractional unit. You know, dollars and cents.  

## floats

Most base 10 decimal numbers don't have a finite representation in base 2.
The canonical example of this is 0.1.

iEEE754 (with bias offset) representation of 0.1 in 64 bits

```
0011 1111 1011 1001 1001 1001 1001 1001 1001 1001 1001 1001 1001 1001 1001 1010
```
note that the mantissa is repeating `1001` and ends with a rounded 1010

Which in base ten is actually something like 
``` 
0.100000001490116119384765625
```

If we add 10 cents and 10 cents together as 0.1, there is no immediate problem, 

```
  0.100000001490116119384765625
+ 0.100000001490116119384765625
-------------------------------
 0.200000003
```

but, eventually

```
  0.100000001490116119384765625
x 512
-------------------------------
0.300000004
```

all the sudden our 

