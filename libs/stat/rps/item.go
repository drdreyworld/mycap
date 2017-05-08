package rps

type Items []Item

func (self *Items) Rotate(n int) {
	if n > 0 {
		capacity := cap(*self)
		values := make([]Item, capacity, capacity)

		if n < capacity {
			copy(values, (*self)[n:])
		}

		*self = values
	}

}

type Item struct {
	Min int64 `json:"min"`
	Max int64 `json:"max"`
	Avg int64 `json:"avg"`
}

func (self *Item) Calc(min int64, max int64, avg int64) {
	if min > 0 {
		if self.Min == 0 || self.Min > min {
			self.Min = min
		}
	}

	if max > 0 {
		if self.Max == 0 || self.Max < max {
			self.Max = max
		}
	}

	if avg > 0 {
		if self.Avg == 0 {
			self.Avg = avg
		} else {
			self.Avg = (self.Avg + avg) / 2
		}
	}
}
