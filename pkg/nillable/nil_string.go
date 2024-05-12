package nillable

import "github.com/golang/protobuf/ptypes/wrappers"

type NilString struct {
	Value   string
	IsEmpty bool
}

func Create(nillableString *wrappers.StringValue) NilString {
	if nillableString == nil {
		return NilString{
			Value:   "",
			IsEmpty: false,
		}
	}
	return NilString{
		Value:   nillableString.GetValue(),
		IsEmpty: true,
	}
}
