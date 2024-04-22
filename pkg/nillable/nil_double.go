package nillable

type NilDouble struct {
	Value   float64
	IsEmpty bool
}

func CreateNillableDouble(val float64) NilDouble {
	if val > 0 {
		return NilDouble{
			Value:   val,
			IsEmpty: false,
		}
	}
	return NilDouble{
		Value:   val,
		IsEmpty: true,
	}
}
