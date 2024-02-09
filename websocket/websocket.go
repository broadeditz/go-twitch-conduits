package websocket

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"

	"github.com/broadeditz/go-twitch-conduits/conduit"
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

		c.handleMessage(data)
	}
}

func (c *Client) handleMessage(data []byte) {
	var message Message
	if err := json.Unmarshal(data, &message); err != nil {
		fmt.Printf("error unmarshaling system message: %+v\nsystem message: %+v\n", err, string(data))
		return
	}

	// TODO: clean this up
	switch message.Metadata.MessageType {
	case MessageTypeNotification:
		switch message.Metadata.SubscriptionType {
		case conduit.EventTypeChannelMessage:
			c.handleChannelMessage(message.Payload)
			return
		default:
			fmt.Printf("unknown notification type: %+v\n", string(data))
		}

	case MessageTypeWelcome:
		var payload SystemMessagePayload
		err := json.Unmarshal(message.Payload, &payload)
		if err != nil {
			fmt.Printf("error unmarshaling welcome message: %+v\n", err)
			return
		}
		if payload.Session == nil {
			fmt.Println("invalid welcome message")
			return
		}

		c.sessionID = payload.Session.ID
		close(c.ready)

	case MessageTypeKeepalive:
	default:
		fmt.Printf("unknown system message type: %+v\n", string(data))
	}
}

func (c *Client) handleChannelMessage(data []byte) {
	message, err := conduit.ParseChannelMessage(data)
	if err != nil {
		fmt.Printf("error: %+v\n", err)
		return
	}

	if c.onChannelMessage != nil {
		c.onChannelMessage(message)
	}
}
