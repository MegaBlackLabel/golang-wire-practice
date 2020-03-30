//+build wireinject

package main

// The build tag makes sure the stub is not built in the final build.
import "github.com/google/wire"

func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}

func InitializeInterface() *MyFooer {
	wire.Build(SetInterface)
	return nil
}

func InitializeStruct() FooBar {
	wire.Build(SetStruct)
	return FooBar{}
}
