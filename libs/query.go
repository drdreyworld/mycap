package libs

import (
	"net"
	"time"

	"github.com/google/gopacket/layers"
)

type Query struct {
	Query string `json:"query"`

	Start time.Time `json:"start"`
	Stop  time.Time `json:"stop"`

	Duration time.Duration `json:"duration"`

	SrcIP net.IP `json:"srcip"`
	DstIP net.IP `json:"dstip"`

	SrcPort layers.TCPPort `json:"srcport"`
	DstPort layers.TCPPort `json:"dstport"`

	ID string `json:"id"`

	ErrorCode     int    `json:"error_code"`
	ErrorMessaget string `json:"error_message"`

	ResponseSize int `json:"response_size"`
}
