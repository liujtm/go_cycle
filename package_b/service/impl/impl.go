package impl

import (
	"aaa/package_i"
	"fmt"
)

func NewB() *PackageBInner {
	return &PackageBInner{}
}

func UpdateB(b *PackageBInner, a package_i.PackageAInterface) *PackageB {
	b.A = a
	return &PackageB{b}
}

type PackageB struct {
	*PackageBInner
}

type PackageBInner struct {
	A package_i.PackageAInterface
}

func (b *PackageBInner) PrintB() {
	fmt.Println("I'm b!")
}

func (b *PackageBInner) PrintAll() {
	fmt.Println("I'm PrintAll from b!")
	b.PrintB()
	b.A.PrintA()
}
