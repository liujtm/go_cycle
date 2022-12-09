使用wire工具遇到循环引用怎么办？

A和B互相依赖对方，可以抽出一层公共的interface(参考 [golang包循环引用的几种解决方案](https://libuba.com/2020/11/02/golang%E5%8C%85%E5%BE%AA%E7%8E%AF%E5%BC%95%E7%94%A8%E7%9A%84%E5%87%A0%E7%A7%8D%E8%A7%A3%E5%86%B3%E6%96%B9%E6%A1%88/))，示例代码如下：

```go
package main
 
import (
	"cycle/package_a"
	"cycle/package_b"
)
 
func main() {
	a := new(package_a.PackageA)
	b := new(package_b.PackageB)
	a.B = b
	b.A = a
	a.PrintAll()
	b.PrintAll()
}
```

如何用[wire依赖注入工具](https://github.com/google/wire)来生成上述代码？核心在于下述的UpdateB函数，通过B生成B自己，但是将interface赋值了。
```go
type PackageB2 package_b.PackageB

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
	return (*PackageB2)(b)
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
```

