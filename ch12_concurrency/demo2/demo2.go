package main

import (
	"errors"
	"fmt"
	"time"
)

func spawn(f func() error) <-chan error {
	c := make(chan error)
	go func() {
		c <- f()
	}()
	// close(c)
	return c
}
func main() {
	c := spawn(func() error {
		time.Sleep(time.Second * 2)
		return errors.New("timeout")
	})
	// c <- errors.New("ssss")
	fmt.Println(<-c)
}
