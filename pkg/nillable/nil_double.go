package nillable

import "github.com/golang/protobuf/ptypes/wrappers"

type NilDouble struct {
	Value   float64
	IsEmpty bool
}

func CreateNillableDouble(nillableInt *wrappers.DoubleValue) NilDouble {
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
