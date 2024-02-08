package websocket

type SystemMessageType string

const (
	SystemMessageTypeWelcome   SystemMessageType = "session_welcome"
	SystemMessageTypeKeepalive SystemMessageType = "session_keepalive"
)

type SystemMessage struct {
	Metadata SystemMessageMetadata `json:"metadata"`
	Payload  SystemMessagePayload  `json:"payload"`
}

type SystemMessageMetadata struct {
	MessageID        string            `json:"message_id"`
	MessageType      SystemMessageType `json:"message_type"`
	MessageTimestamp string            `json:"message_timestamp"`
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
