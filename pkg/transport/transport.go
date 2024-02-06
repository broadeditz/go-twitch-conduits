package transport

type Method string

const (
	MethodWebhook   Method = "webhook"
	MethodWebsocket Method = "websocket"
)

// Transport is used to manage websocket & webhook conduits as a single interface
type Transport interface {
	OnChannelMessage(func(message ChannelMessage))
}
