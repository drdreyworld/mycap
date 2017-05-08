package web

import (
	"bytes"
	"html/template"
	"mycap/libs/agrqueries"
	"mycap/libs/web/paginator"
	"net/http"
	"strconv"
)

func (self *Server) RegisterQueriesActions() {
	http.HandleFunc("/queries/index", self.ActionQueriesIndex)
	http.HandleFunc("/queries/top-by-avg", self.ActionQueriesTopByAvg)
	http.HandleFunc("/queries/top-by-count", self.ActionQueriesTopByCount)
}

func (self *Server) ActionQueriesIndex(w http.ResponseWriter, r *http.Request) {
	content := new(bytes.Buffer)

	page, _ := strconv.Atoi(r.FormValue("page"))

	pgn := paginator.Paginator{}
	pgn.Create(page, 50, self.QueriesCollector.queries.Queries)

	var items []agrqueries.Query

	if pgn.ItemsStop > 0 {
		items = self.QueriesCollector.queries.Queries.Items[pgn.ItemsStart:pgn.ItemsStop]
	}

	self.templates.ExecuteTemplate(content, "views/queries/index", map[string]interface{}{
		"queries":   items,
		"paginator": pgn,
		"tab":       "index",
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
