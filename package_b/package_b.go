package package_b

import (
	"aaa/package_i"
	"fmt"
)

type PackageB struct {
	A package_i.PackageAInterface
}

func (b PackageB) PrintB() {
	fmt.Println("I'm b!")
}

func (b PackageB) PrintAll() {
	b.PrintB()
	b.A.PrintA()
}
