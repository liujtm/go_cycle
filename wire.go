//go:build wireinject
// +build wireinject

package main

import (
	"aaa/package_a"
	ASchema "aaa/package_a/schema"
	"aaa/package_b"
	BSchema "aaa/package_b/schema"

	"github.com/google/wire"
)

type application struct {
	A *ASchema.ASchema
	B *BSchema.BSchema
}

func InitializeApplication() (*application, error) {
	wire.Build(wire.NewSet(
		package_a.ASchemaSet,
		package_b.BSchemaSet,

		wire.Struct(new(application), "*"),
	),
	)
	return &application{}, nil
}
