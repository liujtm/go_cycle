package impl

import (
	"aaa/package_i"
	"fmt"
)

func NewA() *PackageAInner {
	return &PackageAInner{}
}

func UpdateA(a *PackageAInner, b package_i.PackageBInterface) *PackageA {
	a.B = b
	return &PackageA{a}
}

// 匿名字段，PackageA 继承了 PackageAInner 的全部方法
type PackageA struct {
	*PackageAInner
}

type PackageAInner struct {
	B package_i.PackageBInterface
}

func (a *PackageAInner) PrintA() {
	fmt.Println("I'm a!")
}

func (a *PackageAInner) PrintAll() {
	fmt.Println("I'm PrintAll from a!")
	a.PrintA()
	a.B.PrintB()
}
