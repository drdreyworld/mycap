package rps

type Values []int64

func (self *Values) Rotate(n int) {
	if n > 0 {
		capacity := cap(*self)
		values := make([]int64, capacity, capacity)

		if n < capacity {
			copy(values, (*self)[n:])
		}

		*self = values
	}
}
