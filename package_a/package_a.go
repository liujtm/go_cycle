package package_a
 
import (
	"aaa/package_i"
	"fmt"
)
 
type PackageA struct {
	B package_i.PackageBInterface
}
 
func (a PackageA) PrintA() {
	fmt.Println("I'm a!")
}
 
func (a PackageA) PrintAll() {
	a.PrintA()
	a.B.PrintB()
}