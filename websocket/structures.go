package websocket

import (
	"encoding/json"

	"github.com/broadeditz/go-twitch-conduits/conduit"
)

type MessageType string

const (
	MessageTypeWelcome      MessageType = "session_welcome"
	MessageTypeKeepalive    MessageType = "session_keepalive"
	MessageTypeNotification MessageType = "notification"
)

type Message struct {
	Metadata MessageMetadata `json:"metadata"`
	Payload  json.RawMessage `json:"payload"`
}

type MessageMetadata struct {
	MessageID           string            `json:"message_id"`
	MessageType         MessageType       `json:"message_type"`
	MessageTimestamp    string            `json:"message_timestamp"`
	SubscriptionType    conduit.EventType `json:"subscription_type,omitempty"`
	SubscriptionVersion string            `json:"subscription_version,omitempty"`
}

type SystemMessagePayload struct {
	Session *SystemMessagePayloadSession `json:"session,omitempty"`
}

type SystemMessagePayloadSession struct {
	ID                      string `json:"id"`
	Status                  string `json:"status"`
	ConectedAt              string `json:"connected_at"`
	KeepaliveTimeoutSeconds int    `json:"keepalive_timeout_seconds"`
	ReconnectUrl            string `json:"reconnect_url,omitempty"`
}
