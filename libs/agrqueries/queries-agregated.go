package agrqueries

type QueriesAgregated struct {
	Queries Queries           `json:"queries"`
	TopAvg  QueriesTopByAvg   `json:"top_avg"`
	TopCnt  QueriesTopByCount `json:"top_cnt"`
}

func (self *QueriesAgregated) Add(query Query) {
	if len(query.Query) < 1 {
		return
	}

	if pos := self.Queries.Find(query); pos > -1 {
		exists := self.Queries.Items[pos]
		exists.Avg = (exists.Avg + query.Avg) / 2
		exists.Count += query.Count

		if exists.Min > query.Min {
			exists.Min = query.Min
		}

		if exists.Max < query.Max {
			exists.Max = query.Max
		}

		if exists.RSize == 0 {
			exists.RSize = query.RSize
		} else {
			exists.RSize = (exists.RSize + query.RSize) / 2
		}

		self.Queries.Items[pos] = exists

		self.TopAvg.Add(exists)
		self.TopCnt.Add(exists)
	} else {
		self.Queries.Items = append(self.Queries.Items, query)
		self.TopAvg.Add(query)
		self.TopCnt.Add(query)
	}
}
