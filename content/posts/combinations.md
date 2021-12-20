---
title: "combinations"
date: 2021-12-19T13:18:42-06:00
draft: true
---




```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	var number = []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 25, 50, 75, 100}
	m := make(map[int]struct{})

	for _, n := range number {
		m[n] = struct{}{}
	}

	unique := make([]int, 0, len(m))
	for n := range m {
		unique = append(unique, n)
	}
	sort.Ints(unique)

	var result [][]int
	for i, x := range unique {
		for _, y := range unique[i:] {
			result = append(result, []int{x, y})
		}
	}
	for _, r := range result {
		fmt.Printf("%v\n", r)
	}

}
```

