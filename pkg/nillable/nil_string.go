package nillable

type NilString struct {
	Value   string
	IsEmpty bool
}

func Create(val string) NilString {
	if len(val) > 0 {
		return NilString{
			Value:   val,
			IsEmpty: false,
		}
	}
	return NilString{
		Value:   val,
		IsEmpty: true,
	}
}
