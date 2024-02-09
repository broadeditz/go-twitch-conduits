package conduit

import (
	"encoding/json"
)

func ParseChannelMessage(data []byte) (ChannelMessage, error) {
	var msg ChannelMessage
	err := json.Unmarshal(data, &msg)
	return msg, err
}
