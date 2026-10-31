package main

import (
	"errors"
	"fmt"
)

type Gasimli struct {
}

func (d *Gasimli) Random() {
	fmt.Println("snakes")
}

type Snake interface {
	Random()
}

func f() (err error) {
	defer func() {
		err = errors.New("changed by defer")
	}()
	return nil // err = nil, then defer sets err
}

func g() error {
	var err error
	defer func() { err = errors.New("lost") }()
	return err // value copied before defer runs
}

// g() returns nil

// f() returns "changed by defer"

func main() {
	fmt.Println(f())
	var snakes Snake = nil
	var Snake2 *Gasimli = nil
	snakes = Snake2
	fmt.Println(snakes == nil)

}
