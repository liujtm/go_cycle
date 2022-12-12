package package_a

import (
	"aaa/package_a/schema"
	"aaa/package_a/service"
	"aaa/package_a/service/impl"
	"aaa/package_i"

	"github.com/google/wire"
)

var ASchemaSet = wire.NewSet(
	// 下面4行为了生成 ASchema， 但还缺少一个 PackageBInterface 的实现
	impl.NewA,
	impl.UpdateA,
	schema.NewSchema,
	wire.Bind(new(service.AInterface), new(*impl.PackageA)), // 为schema绑定interface实现

	// 为生成B时，提供 PackageAInterface 的实现
	wire.Bind(new(package_i.PackageAInterface), new(*impl.PackageAInner)),
)
