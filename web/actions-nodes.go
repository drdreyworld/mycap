package web

import (
	"bytes"
	"html/template"
	"net/http"
)

func (self *Server) RegisterNodesActions() {
	http.HandleFunc("/nodes/index", self.ActionsNodesIndex)
}

func (self *Server) ActionsNodesIndex(w http.ResponseWriter, r *http.Request) {
	content := new(bytes.Buffer)

	self.templates.ExecuteTemplate(content, "views/nodes/index", map[string]interface{}{
		"nodes": self.AgentsCollector.agents,
	})

	self.templates.ExecuteTemplate(w, "layout/main", map[string]interface{}{
		"pageTitle": "Все агенты",
		"content":   template.HTML(content.String()),
	})
}
