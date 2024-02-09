package helix

// TransportMethod is the method of transport used by the conduit/subscription
type TransportMethod string

const (
	TransportMethodWebhook   TransportMethod = "webhook"
	TransportMethodWebsocket TransportMethod = "websocket"
	TransportMethodConduit   TransportMethod = "conduit"
)

// EventType is the type of event that the conduit/subscription is for
type EventType string

const (
	// EventTypeNull is an empty event type, and therefore the default value of EventType
	EventTypeNull EventType = ""
	//EventTypeChannelMessage is a chat message
	EventTypeChannelMessage EventType = "channel.chat.message"
	// TODO: Add more event types
)
