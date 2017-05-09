package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Server struct {
	Host string `json:"host"`
	Port int    `json:"port"`

	PathTemplates string `json:"path_templates"`
	PathStatic    string `json:"path_static"`

	templates *template.Template

	AgentsCollector  AgentsCollector  `json:"agents_collector"`
	QueriesCollector QueriesCollector `json:"queries_collector"`

	HeadServerHost string `json:"server_host"`
	HeadServerPort int    `json:"server_port"`
}

func (self *Server) StartAgentsCollector() {
	self.AgentsCollector.server = self
	self.AgentsCollector.Collect()
}

func (self *Server) StartQueriesCollector() {
	self.QueriesCollector.server = self
	self.QueriesCollector.Collect()
}

func (self *Server) InitTemplates() {

	funcMap := template.FuncMap{
		"plotRps":         RenderPlotRps,
		"plotRpsAvg":      RenderPlotRpsAvg,
		"plotDurationAvg": RenderPlotDurationAvg,
	}

	tpl := template.New("")
	tpl.Funcs(funcMap)

	if _, err := tpl.ParseGlob(self.PathTemplates + "/layout/*.html"); err != nil {
		log.Fatal(err)
	}

	if _, err := tpl.ParseGlob(self.PathTemplates + "/views/*/*.html"); err != nil {
		log.Fatal(err)
	}

	self.templates = tpl
}

func (self *Server) StartWebServer() {

	self.RegisterDashboardActions()
	self.RegisterNodesActions()
	self.RegisterQueriesActions()
	self.RegisterCounterActions()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(self.PathStatic))))
	http.ListenAndServe(fmt.Sprintf("%s:%d", self.Host, self.Port), nil)
}
