package main

import "fmt"

func main() {
	b := true
	closure := func() error {
		err := fmt.Errorf("something bad happened!!")
		if b {
			return err
		}
		return nil
	}
	b = false

	err := closure()
	fmt.Printf("%v\n", err)
}
