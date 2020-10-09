// +build wireinject

package main

import (
	"github.com/google/wire"
)

func setup(config AppConfig) (application, func(), error) {
	panic(wire.Build(
		graphQLSet,
		newApplication,
	))
}
