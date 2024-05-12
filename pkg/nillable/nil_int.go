package nillable

import "github.com/golang/protobuf/ptypes/wrappers"

type NilInt struct {
	Value   int64
	IsEmpty bool
}

func CreateNillableInt(nillableInt *wrappers.Int64Value) NilInt {
	if nillableInt == nil {
		return NilInt{
			Value:   0,
			IsEmpty: true,
		}
	}
	return NilInt{
		Value:   nillableInt.GetValue(),
		IsEmpty: false,
	}
}
