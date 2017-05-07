package web

import (
	"bytes"
	"html/template"
	"net/http"
)

func (self *Server) RegisterQueriesActions() {
	http.HandleFunc("/queries/index", self.ActionQueriesIndex)
	http.HandleFunc("/queries/top-by-avg", self.ActionQueriesTopByAvg)
	http.HandleFunc("/queries/top-by-count", self.ActionQueriesTopByCount)
}

func (self *Server) ActionQueriesIndex(w http.ResponseWriter, r *http.Request) {
	content := new(bytes.Buffer)

	self.templates.ExecuteTemplate(content, "views/queries/index", map[string]interface{}{
		"queries": self.QueriesCollector.queries.Queries.Items,
		"tab":     "index",
	})

	self.templates.ExecuteTemplate(w, "layout/main", map[string]interface{}{
		"pageTitle": "Все запросы",
		"content":   template.HTML(content.String()),
	})
}

func (self *Server) ActionQueriesTopByAvg(w http.ResponseWriter, r *http.Request) {
	self.QueriesCollector.queries.TopAvg.SortDesc()

	content := new(bytes.Buffer)
	self.templates.ExecuteTemplate(content, "views/queries/top-by-avg", map[string]interface{}{
		"queries": self.QueriesCollector.queries.TopAvg.Items,
		"tab":     "top-by-avg",
	})

	self.templates.ExecuteTemplate(w, "layout/main", map[string]interface{}{
		"pageTitle": "Топ-запросов по времени выполнения",
		"content":   template.HTML(content.String()),
	})
}

func (self *Server) ActionQueriesTopByCount(w http.ResponseWriter, r *http.Request) {
	self.QueriesCollector.queries.TopCnt.SortDesc()

	content := new(bytes.Buffer)
	self.templates.ExecuteTemplate(content, "views/queries/top-by-count", map[string]interface{}{
		"queries": self.QueriesCollector.queries.TopCnt.Items,
		"tab":     "top-by-count",
	})

	self.templates.ExecuteTemplate(w, "layout/main", map[string]interface{}{
		"pageTitle": "Топ-запросов по частоте выполнения",
		"content":   template.HTML(content.String()),
	})
}
