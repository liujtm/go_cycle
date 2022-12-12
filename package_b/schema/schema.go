package schema

import "aaa/package_b/service"

type BSchema struct {
	service.BInterface
}

func NewSchema(B service.BInterface) *BSchema {
	return &BSchema{B}
}
