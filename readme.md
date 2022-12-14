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

如何用[wire依赖注入工具](https://github.com/google/wire)来生成上述代码？核心在于下述的UpdateA函数，通过A生成A自己，但是将interface赋值了。
```go
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

// PackageAMiddle，中间产物(半成品)，NewA()生成，但 PackageBInterface 未绑定实现
type PackageAMiddle struct {
	B package_i.PackageBInterface
}
```

每个package有 PackageAMiddle 和 PackageA 两个版本；PackageAMiddle 给 package_i.PackageAInterface 绑定；PackageA 给自己的schema层绑定；
```go
var ASchemaSet = wire.NewSet(
	// 下面4行为了生成 ASchema， 但还缺少一个 PackageBInterface 的实现
	impl.NewA,
	impl.UpdateA,
	schema.NewSchema,
	wire.Bind(new(service.AInterface), new(*impl.PackageA)), // 为schema绑定interface实现

	// 为生成B时，提供 PackageAInterface 的实现
	wire.Bind(new(package_i.PackageAInterface), new(*impl.PackageAMiddle)),
)
```
最终生成的 wire_gen.go 如下
```go
func InitializeApplication() (*application, error) {
	packageAMiddle := impl.NewA()  // A的半成品、中间产物
	packageBMiddle := impl2.NewB() // B的半成品、中间产物

	packageA := impl.UpdateA(packageAMiddle, packageBMiddle) // 最终可用的A
	aSchema := schema.NewSchema(packageA) // 给schema层使用

	packageB := impl2.UpdateB(packageBMiddle, packageAMiddle) // 最终可用的B
	bSchema := schema2.NewSchema(packageB) // 给schema层使用

	mainApplication := &application{
		A: aSchema,
		B: bSchema,
	}
	return mainApplication, nil
}
```