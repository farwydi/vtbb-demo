package main

import (
	"github.com/farwydi/vtbb-demo/pkg/graphql"
	"github.com/google/wire"
)

var graphQLSet = wire.NewSet(
	graphql.NewResolver,
	graphql.NewRouter,
)
