package helix

import (
	"encoding/json"
)

// ParseChannelMessage parses a channel message from a slice of bytes
func ParseChannelMessage(data []byte) (ChannelMessage, error) {
	var msg ChannelMessage
	err := json.Unmarshal(data, &msg)
	return msg, err
}
