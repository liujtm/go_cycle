package impl

import (
	"aaa/package_i"
	"fmt"
)

func NewB() *PackageBMiddle {
	return &PackageBMiddle{}
}

func UpdateB(b *PackageBMiddle, a package_i.PackageAInterface) *PackageB {
	b.A = a
	return &PackageB{b}
}

type PackageB struct {
	*PackageBMiddle
}

type PackageBMiddle struct {
	A package_i.PackageAInterface
}

func (b *PackageBMiddle) PrintB() {
	fmt.Println("I'm b!")
}

func (b *PackageBMiddle) PrintAll() {
	fmt.Println("I'm PrintAll from b!")
	b.PrintB()
	b.A.PrintA()
}
