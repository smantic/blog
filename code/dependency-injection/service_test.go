package main

import (
	"io"
	"testing"
)

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
