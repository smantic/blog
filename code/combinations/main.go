package main

import (
	"flag"
	"fmt"
	"sort"
)

func main() {

	k := *flag.Int("k", 2, "k sized combination")
	flag.Parse()

	//var number = []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 25, 50, 75, 100}

	nums := []int{1, 2, 3, 4}

	result := Combinations(nums, k)
	for _, r := range result {
		fmt.Printf("%v\n", r)
	}
}

func Combinations(set []int, k int) [][]int {

	m := make(map[int]struct{})

	for _, n := range set {
		m[n] = struct{}{}
	}

	unique := make([]int, 0, len(m))
	for n := range m {
		unique = append(unique, n)
	}
	sort.Ints(unique)

	if k > len(set) {
		return [][]int{set}
	}

	var result [][]int
	for i, _ := range unique {
		for j := range unique[i+1:] {
			sub := make([]int, 0, k)
			for l := 0; l < k; l++ {
				if j+l >= len(unique) {
					break
				}
				sub = append(sub, unique[j+l])
			}
			result = append(result, sub)
		}
	}

	return result
}
