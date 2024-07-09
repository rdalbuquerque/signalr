package signalr

import (
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type TransportType string

var TransportWebSockets TransportType = "WebSockets"
var TransportWebTransports TransportType = "WebTransports"
var TransportServerSentEvents TransportType = "ServerSentEvents"

type TransferFormatType string

var TransferFormatText TransferFormatType = "Text"
var TransferFormatBinary TransferFormatType = "Binary"

type availableTransport struct {
	Transport       string   `json:"transport"`
	TransferFormats []string `json:"transferFormats"`
}

type negotiateResponse struct {
	ConnectionToken     string               `json:"connectionToken,omitempty"`
	ConnectionID        string               `json:"connectionId"`
	NegotiateVersion    int                  `json:"negotiateVersion,omitempty"`
	ProtocolVersion     string               `json:"protocolVersion,omitempty"`
	AvailableTransports []availableTransport `json:"availableTransports,omitempty"`
	TryWebSockets       bool                 `json:"tryWebSockets,omitempty"`
}

// getConnectionVersion returns the maximum value between NegotiateVersion and ProtocolVersion
func (nr *negotiateResponse) getConnectionVersion() (int, error) {
	protocolMajorVersion, err := strconv.Atoi(strings.Split(nr.ProtocolVersion, ".")[0])
	if err != nil {
		return 0, err
	}
	connver := slices.Max([]int{nr.NegotiateVersion, protocolMajorVersion})
	return connver, nil
}

func (nr *negotiateResponse) allowWebSockets() bool {
	return nr.TryWebSockets
}

func (nr *negotiateResponse) hasTransport(transportType TransportType) bool {
	for _, transport := range nr.AvailableTransports {
		if transport.Transport == string(transportType) {
			return true
		}
	}
	return false
}
