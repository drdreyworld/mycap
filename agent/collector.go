package agent

import (
	"fmt"
	"log"
	"mycap/libs"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type Collector struct {
	Device string `json:"device"`

	BPFFilter     string `json:"bpf_filter"`
	BPFFilterPort int    `json:"bpf_filter_port"`
	TXTFilter     string `json:"txt_filter"`

	MaxQueryLen int `json:"max_query_len"`

	buffer  map[string]libs.Query
	queries libs.Queries
}

func (self *Collector) Collect() {

	self.buffer = make(map[string]libs.Query)

	handle, err := pcap.OpenLive(self.Device, int32(self.MaxQueryLen)+5, true, time.Second)
	defer handle.Close()

	if err != nil {
		log.Fatal(err)
	}

	self.BPFFilter = fmt.Sprintf("tcp and port %d", self.BPFFilterPort)
	log.Println("self.BPFFilter", self.BPFFilter)
	handle.SetBPFFilter(self.BPFFilter)

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	var (
		ipLayer  *layers.IPv4
		tcpLayer *layers.TCP
		ok       bool
	)

	for packet := range packetSource.Packets() {
		if applicationLayer := packet.ApplicationLayer(); applicationLayer == nil {
			continue
		} else {
			if ipLayer, ok = packet.Layer(layers.LayerTypeIPv4).(*layers.IPv4); !ok {
				continue
			}
			if tcpLayer, ok = packet.Layer(layers.LayerTypeTCP).(*layers.TCP); !ok {
				continue
			}

			payload := applicationLayer.Payload()

			if !isValidPacket(&payload) {
				continue
			}

			if isQuery(&payload) {
				ID := fmt.Sprintf("%s:%d@%s:%d", ipLayer.SrcIP, tcpLayer.SrcPort, ipLayer.DstIP, tcpLayer.DstPort)
				// log.Println("isQuery", "ID", ID, getQueryText(&payload))

				self.buffer[ID] = libs.Query{

					Query: getQueryText(&payload),
					Start: packet.Metadata().Timestamp,

					SrcIP: ipLayer.SrcIP,
					DstIP: ipLayer.DstIP,

					SrcPort: tcpLayer.SrcPort,
					DstPort: tcpLayer.DstPort,

					ID: ID,
				}
			} else {
				ID := fmt.Sprintf("%s:%d@%s:%d", ipLayer.DstIP, tcpLayer.DstPort, ipLayer.SrcIP, tcpLayer.SrcPort)

				if query, found := self.buffer[ID]; found {

					query.ID = fmt.Sprintf(
						"%s:%d@%s:%d %d\n",
						ipLayer.DstIP, tcpLayer.DstPort, ipLayer.SrcIP, tcpLayer.SrcPort, time.Now().UnixNano(),
					)
					query.Stop = packet.Metadata().Timestamp
					query.Duration = packet.Metadata().Timestamp.Sub(query.Start)

					if isErr(&payload) {
						// log.Println("IsErr", ID, "", query.Query)
						query.ResponseSize = getLength(&payload)
						query.ErrorCode, query.ErrorMessaget = getError(&payload)

						self.queries = append(self.queries, query)
						delete(self.buffer, ID)
					} else if isOk(&payload) || isEOF(&payload) {
						// log.Println("IsOk", ID, "", query.Query)
						query.ResponseSize += (len(payload) - 5)

						self.queries = append(self.queries, query)
						delete(self.buffer, ID)
					} else if getLength(&payload) == 1 {
						// log.Println("Response len = 1", ID, "", query.Query)
						query.ResponseSize += (len(payload) - 5)

						self.queries = append(self.queries, query)
						delete(self.buffer, ID)
					} else {
						log.Println("Not final result", ID, "", query.Query)
						query.ResponseSize += (len(payload) - 5)

						self.buffer[ID] = query
					}
				}
			}
		}
	}
}

type tpayload []byte

func isValidPacket(p *[]byte) bool {
	return len((*p)) >= 5 && getLength(p) > 0
}

func getLength(p *[]byte) int {
	return int((*p)[0]) | int((*p)[1])<<8 | int((*p)[2])<<16
}

func getCode(p *[]byte) byte {
	return (*p)[4]
}

func getSequence(p *[]byte) byte {
	return (*p)[3]
}

func isQuery(p *[]byte) bool {
	return len(*p) == getLength(p)+4 && getCode(p) == 0x03
}

func getQueryText(p *[]byte) string {
	return string((*p)[3+1+1 : 3+1+1+getLength(p)-1])
}

func isOk(p *[]byte) bool {
	return getCode(p) == 0x00
}

func isErr(p *[]byte) bool {
	return getCode(p) == 0xff
}

func isEOF(p *[]byte) bool {
	return getCode(p) == 0xfe
}

func getError(p *[]byte) (code int, message string) {
	code = int((*p)[5]) | int((*p)[6])<<8
	message = string((*p)[3+1+1+2:])
	// if capabilities & CLIENT_PROTOCOL_41 {
	// https://dev.mysql.com/doc/dev/mysql-server/latest/page_protocol_basic_err_packet.html
	// string[1] sql_state_marker 		# marker of the SQL state
	// string[5] sql_state 						SQL state
	// message = string((*p)[3+1+1+2   + 1 + 5:])
	// }

	return code, message
}
