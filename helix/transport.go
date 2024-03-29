package helix

// Transport is used to manage websocket & webhook conduits as a single interface
type Transport interface {
	// Init initializes the transport
	Init() error
	// Close closes the transport
	Close()
	// Ready returns the ready channel, which is closed when the transport is ready to be used.
	// Reading from the channel, after calling Init, will block until the transport is ready.
	Ready() chan struct{}
	// GetTransportUpdate returns a TransportUpdate used to update the conduit
	GetTransportUpdate() *TransportUpdate
	// OnChannelMessage sets a callback to be called when a channel message is received, is not executed in gorourines. It is the responsibility of the caller to handle concurrency.
	OnChannelMessage(func(message ChannelMessage))
	// TODO: Add more event types
}
