package duration

type Values []float64

func (self *Values) Rotate(n int) {
	if n > 0 {
		capacity := cap(*self)
		values := make(Values, capacity, capacity)

		if n < capacity {
			copy(values, (*self)[n:])
		}

		*self = values
	}
}
