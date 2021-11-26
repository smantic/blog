package panic

import (
	"testing"
)

var i interface{}

func BenchmarkPanic8Frames(b *testing.B) {
	for n := 0; n < b.N; n++ {
		done := make(chan struct{})
		go func() {
			defer func() {
				i = recover()
				done <- struct{}{}
				close(done)
			}()
			func() {
				func() {
					func() {
						func() {
							func() {
								func() {
									func() {
										func() {
											panic("panic!!")
										}()
									}()

								}()
							}()

						}()
					}()

				}()
			}()
		}()
		<-done
	}
}

func BenchmarkPanic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		done := make(chan struct{})
		go func() {
			defer func() {
				i = recover()
				done <- struct{}{}
				close(done)
			}()
			panic("panic!!")
		}()
		<-done
	}
}

func BenchmarkBase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		done := make(chan struct{})
		go func() {
			done <- struct{}{}
			close(done)
		}()
		<-done
	}
}
