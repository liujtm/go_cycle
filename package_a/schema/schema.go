package schema

import "aaa/package_a/service"

type ASchema struct {
	service.AInterface
}

func NewSchema(A service.AInterface) *ASchema {
	return &ASchema{A}
}
