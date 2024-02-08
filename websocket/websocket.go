package websocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"

	conduit2 "github.com/broadeditz/go-twitch-conduits/conduit"
)

func (c *Client) Init() error {
	conn, _, err := websocket.DefaultDialer.Dial("wss://eventsub.wss.twitch.tv/ws", nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	c.conn = conn

	done := make(chan struct{})

	// Main read loop for the websocket connection
	go c.readMessages(done)

	select {
	case <-done:
		return nil

	case <-c.interrupt:
		err = c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			return err
		}

		select {
		case <-done:
		case <-time.After(time.Second):
		}

		return nil
	}
}

func (c *Client) readMessages(done chan struct{}) {
	defer close(done)
	for {
		_, data, err := c.conn.ReadMessage()
		if err != nil {
			return
		}

		// TODO: remove debug logging
		fmt.Printf("%+v\n", string(data))

		reader := bytes.NewReader(data)
		messageType, err := conduit2.ParseMessageType(reader)
		if err != nil {
			fmt.Printf("websocket parse error: %+v\n", err)
			continue
		}

		switch messageType {
		case conduit2.EventTypeChannelMessage:
			c.handleChannelMessage(data)

			// Null EventType is used for non-subscription messages, meaning system messages mostly
		case conduit2.EventTypeNull:
			c.handleSystemMessage(data)

			// Unknown EventType means message type is not yet implemented
		case conduit2.EventTypeUnknown:
			fmt.Printf("unknown message type: %+v\n", string(data))
		}
	}
}

func (c *Client) handleChannelMessage(data []byte) {
	message, err := conduit2.ParseChannelMessage(data)
	if err != nil {
		fmt.Printf("error: %+v\n", err)
		return
	}

	if c.onChannelMessage != nil {
		c.onChannelMessage(message)
	}
}

func (c *Client) handleSystemMessage(data []byte) {
	var message SystemMessage
	if err := json.Unmarshal(data, &message); err != nil {
		fmt.Printf("error unmarshaling system message: %+v\nsystem message: %+v\n", err, string(data))
		return
	}

	switch message.Metadata.MessageType {
	case SystemMessageTypeWelcome:
		if message.Payload.Session == nil {
			fmt.Println("invalid welcome message")
			return
		}

		c.sessionID = message.Payload.Session.ID
		close(c.ready)

	case SystemMessageTypeKeepalive:
	default:
		fmt.Printf("unknown system message type: %+v\n", string(data))
	}
}

func (c *Client) GetTransportUpdate() *conduit2.TransportUpdate {
	return &conduit2.TransportUpdate{
		Method:    conduit2.TransportMethodWebsocket,
		SessionID: c.sessionID,
	}
}

func (c *Client) Close() {
	close(c.interrupt)
}
