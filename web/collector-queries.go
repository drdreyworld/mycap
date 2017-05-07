package web

import (
	"mycap/libs/agrqueries"
	"mycap/libs/client"
	"mycap/libs/stat"
	"time"
)

type QueriesCollector struct {
	server  *Server
	queries agrqueries.QueriesAgregated
	stat    stat.Stat
}

func (self *QueriesCollector) Collect() {
	cli := client.ServerClient{}
	cli.Host = self.server.HeadServerHost
	cli.Port = self.server.HeadServerPort

	for {
		func() {
			response, err := cli.GetQueries()

			if err == nil && response.Error.Code == 0 {
				self.queries = response.Result
			}
		}()

		func() {
			response, err := cli.GetStat()

			if err == nil && response.Error.Code == 0 {
				self.stat = response.Result
			}
		}()

		time.Sleep(time.Second)
	}
}
