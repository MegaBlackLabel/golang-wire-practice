package main

import (
	"fmt"

	"github.com/google/wire"
)

type Message string

type Greeter struct {
	Message Message
}

type Event struct {
	Greeter Greeter
}

func NewMessage() Message {
	return Message("hello wire")
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (g Greeter) Greet() Message {
	return g.Message
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

type Fooer interface {
	Foo() string
}

type MyFooer string

func (b *MyFooer) Foo() string {
	return string(*b)
}

func provideMyFooer() *MyFooer {
	b := new(MyFooer)
	*b = "Hello, World!"
	return b
}

type Bar string

func provideBar(f Fooer) string {
	// f will be a *MyFooer.
	return f.Foo()
}

var appSet = wire.NewSet(
	provideMyFooer,
	wire.Bind(new(Fooer), new(*MyFooer)),
	provideBar)
