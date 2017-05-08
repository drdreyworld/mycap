package rps

import "time"

type Counter struct {
	Values Values `json:"values"`
	Hour   Items  `json:"hour"`
	Day    Items  `json:"day"`
	Month  Items  `json:"month"`
	Time   int64  `json:"time"`
}

func (self *Counter) Init() {
	self.Values = make(Values, 60, 60)
	self.Hour = make(Items, 60, 60)
	self.Day = make(Items, 24, 24)
	self.Month = make(Items, 30, 30)

	go func() {
		for {
			self.Tick()
			time.Sleep(time.Second)
		}
	}()
}

func (self *Counter) Inc(eventTime int64, value int64) {
	index := int(time.Now().Unix() - eventTime)

	len_values := len(self.Values)

	if index > 0 && index < len_values {

		self.Values[len_values-index] += value

		self.Hour[len(self.Hour)-1].Calc(
			self.Values[len_values-index],
			self.Values[len_values-index],
			self.Values[len_values-index],
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
	if self.Time == 0 {
		self.Time = time.Now().Unix() - 1
	}

	sleep := time.Now().Unix() - self.Time
	self.Time = time.Now().Unix()

	self.Values.Rotate(int(sleep))
	self.Hour.Rotate(int(sleep / 60))
	self.Day.Rotate(int(sleep / 3600))
	self.Month.Rotate(int(sleep / 86400))

	if time.Now().Second() == 0 {
		self.Hour.Rotate(1)
	}

	if time.Now().Minute() == 0 {
		self.Day.Rotate(1)
	}

	if time.Now().Hour() == 0 {
		self.Month.Rotate(1)
	}
}
