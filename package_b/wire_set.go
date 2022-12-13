package package_b

import (
	"aaa/package_b/schema"
	"aaa/package_b/service"
	"aaa/package_b/service/impl"
	"aaa/package_i"

	"github.com/google/wire"
)

var BSchemaSet = wire.NewSet(
	// 下面4行为了生成 BSchema， 但还缺少一个 PackageAInterface 的实现
	impl.NewB,
	impl.UpdateB,
	schema.NewSchema,
	wire.Bind(new(service.BInterface), new(*impl.PackageB)), // 为schema绑定interface实现

	// 为生成A时，提供 PackageBInterface 的实现
	wire.Bind(new(package_i.PackageBInterface), new(*impl.PackageBMiddle)),
)
