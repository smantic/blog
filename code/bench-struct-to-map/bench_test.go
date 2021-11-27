package jsonhack

import "testing"

var foo Foo = Foo{
	A: Bar{"A"},
	B: Bar{"B"},
	C: Bar{"C"},
	D: Bar{"D"},
	E: Bar{"E"},
	F: -1,
	G: 1,
	H: "H",
}

// to elim some compiler optimizations
var result map[string]interface{}

func BenchmarkBase(b *testing.B) {
	var m map[string]interface{}
	for n := 0; n < b.N; n++ {
		m = base(foo)
	}
	result = m
}

func BenchmarkHack(b *testing.B) {
	var m map[string]interface{}
	for n := 0; n < b.N; n++ {
		m = hack(foo)
	}
	result = m
}

func BenchmarkReflect(b *testing.B) {
	var m map[string]interface{}
	for n := 0; n < b.N; n++ {
		m = reflec(foo)
	}
	result = m
}

func BenchmarkMitchellh(b *testing.B) {
	var m map[string]interface{}
	for n := 0; n < b.N; n++ {
		m = mitchellh(foo)
	}
	result = m
}
