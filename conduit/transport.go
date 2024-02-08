package conduit

// Transport is used to manage websocket & webhook conduits as a single interface
type Transport interface {
	// Init initializes the transport
	Init() error
	// Close closes the transport
	Close()
	// Ready return the ready channel
	Ready() chan struct{}
	// GetTransportUpdate returns a transport update used to update the conduit
	GetTransportUpdate() *TransportUpdate
	// OnChannelMessage defines the callback called when a channel message is received
	OnChannelMessage(func(message ChannelMessage))
	// TODO: Add more event types
}
