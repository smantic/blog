package main

import "io"

// Foo is something that we can write to.
type Foo interface {
	io.Writer
}

type Service struct {
	f Foo
}

func (s *Service) Bar() error {
	_, err := s.f.Write(nil)
	return err
}
