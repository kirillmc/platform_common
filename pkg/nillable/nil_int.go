package nillable

import (
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type NilInt struct {
	Value   int64
	IsEmpty bool
}

func CreateNillableInt(nillableInt *wrapperspb.Int64Value) NilInt {
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
