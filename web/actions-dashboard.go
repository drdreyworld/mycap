package web

import (
	"bytes"
	"html/template"
	"net/http"
)

func (self *Server) RegisterDashboardActions() {
	http.HandleFunc("/", self.ActionDashboard)
}

func (self *Server) ActionDashboard(w http.ResponseWriter, r *http.Request) {
	content := new(bytes.Buffer)

	self.templates.ExecuteTemplate(content, "views/dashboard/index", map[string]interface{}{
		"stat": self.QueriesCollector.stat,
	})

	self.templates.ExecuteTemplate(w, "layout/main", map[string]interface{}{
		"content": template.HTML(content.String()),
	})
}
