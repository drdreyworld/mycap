package agrqueries

import "sort"

type QueriesTopByCount struct {
	Queries
	MaxItems int `json:"max_items"`
}

func (self *QueriesTopByCount) SortAsc() {
	sort.Sort(self)
}

func (self *QueriesTopByCount) SortDesc() {
	sort.Sort(sort.Reverse(self))
}

func (self *QueriesTopByCount) Add(query Query) {
	if i := self.Find(query); i != -1 {
		self.Items[i] = query
	} else if self.MaxItems > 0 && self.Len() < self.MaxItems {
		self.Items = append(self.Items, query)
	} else if self.Len() == 0 {
		self.Items = append(self.Items, query)
	} else {
		self.SortAsc()
		if self.Items[0].Count < query.Count {
			self.Items[0] = query
		}
	}
}

func (self QueriesTopByCount) Less(i, j int) bool {
	return self.Items[i].Count < self.Items[j].Count
}
