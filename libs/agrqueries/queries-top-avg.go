package agrqueries

import "sort"

type QueriesTopByAvg struct {
	Queries
	MaxItems int `json:"max_items"`
}

func (self *QueriesTopByAvg) SortAsc() {
	sort.Sort(self)
}

func (self *QueriesTopByAvg) SortDesc() {
	sort.Sort(sort.Reverse(self))
}

func (self *QueriesTopByAvg) Add(query Query) {
	if i := self.Find(query); i != -1 {
		self.Items[i] = query
	} else if self.MaxItems > 0 && self.Len() < self.MaxItems {
		self.Items = append(self.Items, query)
	} else if self.Len() == 0 {
		self.Items = append(self.Items, query)
	} else {
		self.SortAsc()
		if self.Items[0].Max < query.Max {
			self.Items[0] = query
		}
	}
}

func (self QueriesTopByAvg) Less(i, j int) bool {
	return self.Items[i].Max < self.Items[j].Max
}
