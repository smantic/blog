package fetchonce

import (
	"sync"
	"time"
)

type Foo struct {
	data *string
	once sync.Once
}

func (f *Foo) GenChannel() <-chan string {

	channel := make(chan string)

	go func() {
		f.once.Do(func() {
			d := fetch()
			f.data = &d
		})
		channel <- *f.data
	}()
	return channel
}

func (f *Foo) Wait() string {

	f.once.Do(func() {
		d := fetch()
		f.data = &d
	})

	return *f.data
}

func fetch() string {
	time.Sleep(time.Millisecond * 500)
	return "imporant data"
}
