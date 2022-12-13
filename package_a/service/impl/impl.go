package impl

import (
	"aaa/package_i"
	"fmt"
)

// 生成中间产物 PackageAMiddle
func NewA() *PackageAMiddle {
	return &PackageAMiddle{}
}

// 通过中间产物 PackageAMiddle 和 package B 绑定的实现， 生成最终的 PackageA
func UpdateA(a *PackageAMiddle, b package_i.PackageBInterface) *PackageA {
	a.B = b
	return &PackageA{a}
}

// 匿名字段，PackageA 继承了 PackageAMiddle 的全部方法
type PackageA struct {
	*PackageAMiddle
}

// PackageAMiddle，中间产物，NewA()生成，但 PackageBInterface 未绑定实现
type PackageAMiddle struct {
	B package_i.PackageBInterface
}

func (a *PackageAMiddle) PrintA() {
	fmt.Println("I'm a!")
}

func (a *PackageAMiddle) PrintAll() {
	fmt.Println("I'm PrintAll from a!")
	a.PrintA()
	a.B.PrintB()
}
