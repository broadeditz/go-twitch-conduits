package conduit

type TransportMethod string

const (
	TransportMethodWebhook   TransportMethod = "webhook"
	TransportMethodWebsocket TransportMethod = "websocket"
)

type EventType string

const (
	EventTypeNull           EventType = ""
	EventTypeUnknown        EventType = "unknown"
	EventTypeChannelMessage EventType = "channel.chat.message"
	// TODO: Add more event types
)

func ParseEventType(s string) EventType {
	switch s {
	case "channel.chat.message":
		return EventTypeChannelMessage
	default:
		return EventTypeUnknown
	}
}
