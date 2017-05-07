package countpertime

import (
	"time"
)

type Counter struct {
	Items []int64 `json:"items"`

	StepSize  int64 `json:"step_size"`  // количество секунд, за которое агрегируется количество
	StepCount int64 `json:"step_count"` // количество шагов, которое хранит счетчик
	TimeZero  int64 `json:"time_zero"`  // unixtime нулевого шага (автоматически NOW())
	TimeMax   int64 `json:"time_max"`   // unixtime max + 1 шага (автоматически TimeZero + StepSize * StepCount)
}

func (self *Counter) IndexByTime(unixTime int64) int {
	return int((unixTime - self.TimeZero) / self.StepSize)
}

func (self *Counter) Init(stepSize int64, stepCount int64) {
	self.StepSize = stepSize
	self.StepCount = stepCount

	self.TimeZero = time.Now().Unix() - self.StepCount*self.StepSize
	self.TimeZero = self.StepSize * (self.TimeZero / self.StepSize)

	self.TimeMax = self.TimeZero + (self.StepSize * self.StepCount)

	self.Items = make([]int64, self.StepCount+1, self.StepCount+1)

	go func() {
		for {
			nowtime := time.Now().Unix()
			if nowtime > self.TimeMax {
				// например вылезли на два шага
				offset := (nowtime - self.TimeMax) / self.StepSize // Ceil ?
				self.TimeZero += offset * self.StepSize
				self.TimeMax += offset * self.StepSize
				// немного бредово
				for i := 0; i < int(offset); i++ {
					self.Items = append(self.Items[1:], 0)
				}
			}
			time.Sleep(time.Second * time.Duration(self.StepSize))
		}
	}()
}

func (self *Counter) Inc(eventTime int64, incSize int64) {
	eventIndex := (eventTime - self.TimeZero) / self.StepSize

	if eventIndex < 0 || eventIndex > self.StepCount {
		return
	}

	self.Items[eventIndex] += incSize

	if eventTime > self.TimeMax {
		self.Items = append(self.Items[1:], 0)
		self.TimeZero += self.StepSize
		self.TimeMax += self.StepSize
	}
}

type Counters struct {
	PerSec  Counter `json:"sec"`
	PerMin  Counter `json:"mins"`
	PerHour Counter `json:"hours"`
	PerDay  Counter `json:"days"`
}

func (self *Counters) Init() {
	self.PerSec.Init(1, 60)
	self.PerMin.Init(60, 60)
	self.PerHour.Init(3600, 24)
	self.PerDay.Init(3600*24, 30)
}

func (self *Counters) Inc(eventTime int64, incSize int64) {
	self.PerSec.Inc(eventTime, incSize)
	self.PerMin.Inc(eventTime, incSize)
	self.PerHour.Inc(eventTime, incSize)
	self.PerDay.Inc(eventTime, incSize)
}
