package web

import "net/http"

func (self *Server) RegisterCounterActions() {
	http.HandleFunc("/counters/rps-sec", self.ActionCounterRpsSecAjax)
	http.HandleFunc("/counters/rps-min", self.ActionCounterRpsMinAjax)
	http.HandleFunc("/counters/rps-hour", self.ActionCounterRpsHourAjax)
	http.HandleFunc("/counters/rps-day", self.ActionCounterRpsDayAjax)
	http.HandleFunc("/counters/duration-sec", self.ActionCounterDurationSecAjax)
	http.HandleFunc("/counters/duration-min", self.ActionCounterDurationMinAjax)
	http.HandleFunc("/counters/duration-hour", self.ActionCounterDurationHourAjax)
	http.HandleFunc("/counters/duration-day", self.ActionCounterDurationDayAjax)
}

func (self *Server) ActionCounterRpsSecAjax(w http.ResponseWriter, r *http.Request) {
	self.templates.ExecuteTemplate(w, "views/counter/rps-sec", map[string]interface{}{
		"stat": self.QueriesCollector.stat,
	})
}

func (self *Server) ActionCounterRpsMinAjax(w http.ResponseWriter, r *http.Request) {
	self.templates.ExecuteTemplate(w, "views/counter/rps-min", map[string]interface{}{
		"stat": self.QueriesCollector.stat,
	})
}

func (self *Server) ActionCounterRpsHourAjax(w http.ResponseWriter, r *http.Request) {
	self.templates.ExecuteTemplate(w, "views/counter/rps-hour", map[string]interface{}{
		"stat": self.QueriesCollector.stat,
	})
}

func (self *Server) ActionCounterRpsDayAjax(w http.ResponseWriter, r *http.Request) {
	self.templates.ExecuteTemplate(w, "views/counter/rps-day", map[string]interface{}{
		"stat": self.QueriesCollector.stat,
	})
}

func (self *Server) ActionCounterDurationSecAjax(w http.ResponseWriter, r *http.Request) {
	self.templates.ExecuteTemplate(w, "views/counter/duration-sec", map[string]interface{}{
		"stat": self.QueriesCollector.stat,
	})
}

func (self *Server) ActionCounterDurationMinAjax(w http.ResponseWriter, r *http.Request) {
	self.templates.ExecuteTemplate(w, "views/counter/duration-min", map[string]interface{}{
		"stat": self.QueriesCollector.stat,
	})
}

func (self *Server) ActionCounterDurationHourAjax(w http.ResponseWriter, r *http.Request) {
	self.templates.ExecuteTemplate(w, "views/counter/duration-hour", map[string]interface{}{
		"stat": self.QueriesCollector.stat,
	})
}

func (self *Server) ActionCounterDurationDayAjax(w http.ResponseWriter, r *http.Request) {
	self.templates.ExecuteTemplate(w, "views/counter/duration-day", map[string]interface{}{
		"stat": self.QueriesCollector.stat,
	})
}
