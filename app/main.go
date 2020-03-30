package main

import "fmt"

func main() {
	e := InitializeEvent()

	d := InitializeFoo()
	a := d.Foo()
	fmt.Println(a)

	e.Start()
	// e.Foo()
}
