//+build wireinject

package main

// The build tag makes sure the stub is not built in the final build.
import "github.com/google/wire"

func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}

func InitializeFoo() *MyFooer {
	wire.Build(appSet)
	return nil
}
