package conduit

type TransportMethod string

const (
	TransportMethodWebhook   TransportMethod = "webhook"
	TransportMethodWebsocket TransportMethod = "websocket"
)

type EventType string

const (
	EventTypeChannelMessage EventType = "channel.chat.message"
)
