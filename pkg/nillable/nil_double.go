package nillable

import (
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type NilDouble struct {
	Value   float64
	IsEmpty bool
}

func CreateNillableDouble(nillableInt *wrapperspb.DoubleValue) NilDouble {
	if nillableInt == nil {
		return NilDouble{
			Value:   0,
			IsEmpty: true,
		}
	}
	return NilDouble{
		Value:   nillableInt.GetValue(),
		IsEmpty: false,
	}
}
