package rps

import "time"

type Counter struct {
	Values []int64 `json:"values"`
	Hour   Items   `json:"hour"`
	Day    Items   `json:"day"`
	Month  Items   `json:"month"`
}

func (self *Counter) Init() {
	self.Values = make([]int64, 60, 60)
	self.Hour = make([]Item, 60, 60)
	self.Day = make([]Item, 24, 24)
	self.Month = make([]Item, 30, 30)

	go func() {
		for {
			self.Tick()
			time.Sleep(time.Second)
		}
	}()
}

func (self *Counter) Inc(eventTime int64, value int64) {
	index := int(time.Now().Unix() - eventTime)

	if index > 0 && index < len(self.Values) {

		self.Values[60-index] += value

		self.Hour[len(self.Hour)-1].Calc(
			self.Values[60-index],
			self.Values[60-index],
			self.Values[60-index],
		)

		self.Day[len(self.Day)-1].Calc(
			self.Hour[len(self.Hour)-1].Min,
			self.Hour[len(self.Hour)-1].Max,
			self.Hour[len(self.Hour)-1].Avg,
		)

		self.Month[len(self.Month)-1].Calc(
			self.Day[len(self.Day)-1].Min,
			self.Day[len(self.Day)-1].Max,
			self.Day[len(self.Day)-1].Avg,
		)
	}
}

func (self *Counter) Tick() {
	self.Values = append(self.Values[1:], 0)

	if time.Now().Second() == 0 {
		self.Hour.Rotate()
	}

	if time.Now().Minute() == 0 {
		self.Hour.Rotate()
	}

	if time.Now().Hour() == 0 {
		self.Hour.Rotate()
	}
}
