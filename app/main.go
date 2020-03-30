package main

import "fmt"

func main() {
	e := InitializeEvent()
	e.Start()

	d := InitializeInterface()
	fmt.Println(d.Foo())

	s := InitializeStruct()
	fmt.Println(s.MyBar)
	fmt.Println(s.MyFoo)
}
