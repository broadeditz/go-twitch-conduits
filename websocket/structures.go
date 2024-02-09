package websocket

import (
	"encoding/json"

	"github.com/broadeditz/go-twitch-conduits/helix"
)

// MessageType is the type of message received from the websocket
type MessageType string

const (
	MessageTypeWelcome      MessageType = "session_welcome"
	MessageTypeKeepalive    MessageType = "session_keepalive"
	MessageTypeNotification MessageType = "notification"
)

// Message is the structure of a message received from the websocket
type Message struct {
	Metadata MessageMetadata `json:"metadata"`
	Payload  json.RawMessage `json:"payload"`
}

// MessageMetadata is the metadata of a message received from the websocket
type MessageMetadata struct {
	MessageID           string          `json:"message_id"`
	MessageType         MessageType     `json:"message_type"`
	MessageTimestamp    string          `json:"message_timestamp"`
	SubscriptionType    helix.EventType `json:"subscription_type,omitempty"`
	SubscriptionVersion string          `json:"subscription_version,omitempty"`
}

// SystemMessagePayload is the payload of a system/session message
type SystemMessagePayload struct {
	Session *SystemMessagePayloadSession `json:"session,omitempty"`
}

// SystemMessagePayloadSession is the payload of a system/session message
type SystemMessagePayloadSession struct {
	ID                      string `json:"id"`
	Status                  string `json:"status"`
	ConectedAt              string `json:"connected_at"`
	KeepaliveTimeoutSeconds int    `json:"keepalive_timeout_seconds"`
	ReconnectUrl            string `json:"reconnect_url,omitempty"`
}
