package agrqueries

import (
	"crypto/md5"
	"io"
	"time"
)

type Query struct {
	Query string
	From  string
	To    string
	Hash  string
	Avg   time.Duration
	Min   time.Duration
	Max   time.Duration
	Count int
	RSize int
}

func (self *Query) GetQuery() string {
	return self.Query
}

func (self *Query) CalcHash() {
	h := md5.New()
	io.WriteString(h, self.Query)
	io.WriteString(h, self.From)
	io.WriteString(h, self.To)

	self.Hash = string(h.Sum(nil))
}

func (self *Query) GetHash() string {
	return self.Hash
}

func CreateQuery(from string, to string, query string, rsize int, duration time.Duration) Query {
	result := Query{
		Query: query,
		From:  from,
		To:    to,
		Avg:   duration,
		Min:   duration,
		Max:   duration,
		Count: 1,
		RSize: rsize,
	}
	result.CalcHash()
	return result
}
