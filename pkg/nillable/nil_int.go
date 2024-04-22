package nillable

type NilInt struct {
	Value   int64
	IsEmpty bool
}

func CreateNillableInt(val int64) NilInt {
	if val > 0 {
		return NilInt{
			Value:   val,
			IsEmpty: false,
		}
	}
	return NilInt{
		Value:   val,
		IsEmpty: true,
	}
}
