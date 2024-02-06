package transport

// Transport is used to manage websocket & webhook conduits as a single interface
type Transport interface {
	OnChannelMessage(func(message ChannelMessage))
}
