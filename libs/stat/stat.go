package stat

import (
	"mycap/libs"
	"mycap/libs/stat/duration"
	"mycap/libs/stat/rps"
)

type Stat struct {
	Rps      rps.Counter      `json:"rps"`
	Duration duration.Counter `json:"duration"`
}

func (self *Stat) Init() {
	self.Rps.Init()
	self.Duration.Init()
}

func (self *Stat) FixInfo(query libs.Query) {
	self.Rps.Inc(query.Start.Unix(), 1)
	self.Duration.Inc(query.Start.Unix(), query.Duration.Seconds())
}
