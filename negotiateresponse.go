package signalr

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
	AvailableTransports []availableTransport `json:"availableTransports,omitempty"`
	TryWebSockets       bool                 `json:"tryWebSockets,omitempty"`
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
