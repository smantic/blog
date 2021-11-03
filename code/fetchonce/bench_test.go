package fetchonce

import "testing"

var result string

func BenchmarkBlocking(b *testing.B) {
	var foo Foo
	for n := 0; n < b.N; n++ {
		a := foo.Wait()
		result = a
	}
}

func BenchmarkChan(b *testing.B) {
	var foo Foo
	for n := 0; n < b.N; n++ {
		ch := foo.GenChannel()
		result = <-ch
	}
}
