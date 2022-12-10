//go:build wireinject
// +build wireinject

package main

import (
	"aaa/package_a"
	"aaa/package_b"
	"aaa/package_i"

	"github.com/google/wire"
)

// 匿名字段，PackageB2继承了PackageB的全部方法
type PackageB2 struct {
	*package_b.PackageB
}

type application struct {
	A *package_a.PackageA
	B *PackageB2
}

func NewA(i package_i.PackageBInterface) *package_a.PackageA {
	a := new(package_a.PackageA)
	a.B = i
	return a
}

func NewB() *package_b.PackageB {
	return new(package_b.PackageB)
}

func UpdateB(b *package_b.PackageB, i package_i.PackageAInterface) *PackageB2 {
	b.A = i
	return &PackageB2{b}
}

func InitializeApplication() (*application, error) {
	wire.Build(wire.NewSet(
		wire.Bind(new(package_i.PackageBInterface), new(*package_b.PackageB)),
		wire.Bind(new(package_i.PackageAInterface), new(*package_a.PackageA)),
		NewA,
		NewB,
		UpdateB,
		wire.Struct(new(application), "*"),
	),
	)
	return &application{}, nil
}
