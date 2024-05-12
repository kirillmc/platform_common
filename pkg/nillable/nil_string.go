package nillable

import (
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type NilString struct {
	Value   string
	IsEmpty bool
}

func Create(nillableString *wrapperspb.StringValue) NilString {
	if nillableString == nil {
		return NilString{
			Value:   "",
			IsEmpty: true,
		}
	}
	return NilString{
		Value:   nillableString.GetValue(),
		IsEmpty: false,
	}
}
