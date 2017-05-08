package stat

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mycap/libs"
	"mycap/libs/stat/duration"
	"mycap/libs/stat/rps"
	"time"
)

type Stat struct {
	Rps      rps.Counter      `json:"rps"`
	Duration duration.Counter `json:"duration"`
	DataFile string           `json:"datafile"`
}

func (self *Stat) Init() {
	self.Rps.Init()
	self.Duration.Init()

	if len(self.DataFile) > 0 {
		self.LoadFromFile(self.DataFile)

		go func() {
			for {
				self.SaveToFile(self.DataFile)
				time.Sleep(3 * time.Second)
			}
		}()
	}
}

func (self *Stat) FixInfo(query libs.Query) {
	self.Rps.Inc(query.Start.Unix(), 1)
	self.Duration.Inc(query.Start.Unix(), query.Duration.Seconds())
}

func (self *Stat) SaveToFile(filename string) error {
	content, err := json.Marshal(*self)

	if err != nil {
		log.Println(err)
		return err
	}

	if err = ioutil.WriteFile(filename, content, 0644); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (self *Stat) LoadFromFile(filename string) error {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Println(err)
		return err
	}

	if err := json.Unmarshal(content, self); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
