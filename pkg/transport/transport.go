package transport

// Transport is used to manage websocket & webhook conduits as a single interface
type Transport interface {
	Start() error
	OnChannelMessage(func(message ChannelMessage))
	// TODO: Add more event types
}
