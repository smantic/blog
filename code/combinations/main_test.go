package main

import (
	"testing"

	"github.com/go-test/deep"
)

func TestCombos(t *testing.T) {

	expected := [][]int{
		{1, 2, 3},
		{1, 2, 4},
		{1, 3, 4},
		{2, 3, 4},
	}

	got := Combinations([]int{1, 2, 3, 4}, 3)
	diff := deep.Equal(got, expected)
	if len(diff) > 0 {
		t.Error(diff)
	}
}
