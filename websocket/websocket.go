package websocket

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/broadeditz/go-twitch-conduits/helix"

	"github.com/gorilla/websocket"
)

// Init starts the websocket connection and sets up the read loop
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

// readMessages continuously reads messages from the websocket connection
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

// handleMessage unmarshals the message metadata and handles it based on the message type
func (c *Client) handleMessage(data []byte) {
	var message Message
	if err := json.Unmarshal(data, &message); err != nil {
		fmt.Printf("error unmarshaling system message: %+v\nsystem message: %+v\n", err, string(data))
		return
	}

	switch message.Metadata.MessageType {
	// Notifications are subscription messages
	case MessageTypeNotification:
		c.handleNotificationMessage(message)
	// Welcome messages contain the session ID, and mean the websocket is ready
	case MessageTypeWelcome:
		c.handleWelcomeMessage(message.Payload)

	case MessageTypeKeepalive:
		//TODO: close on keepalive timeout
	default:
		fmt.Printf("unknown system message type: %+v\n", string(data))
	}
}

// handle subscription notifications
func (c *Client) handleNotificationMessage(message Message) {
	switch message.Metadata.SubscriptionType {
	case helix.EventTypeChannelMessage:
		c.handleChannelMessage(message.Payload)

	default:
		fmt.Printf("unknown notification type: %+v\n", message.Metadata.SubscriptionType)
	}
}

// handle chat messages
func (c *Client) handleChannelMessage(data []byte) {
	message, err := helix.ParseChannelMessage(data)
	if err != nil {
		fmt.Printf("error: %+v\n", err)
		return
	}

	if c.onChannelMessage != nil {
		c.onChannelMessage(message)
	}
}

// Welcome message requires us to set the session ID & mark the transport as ready
func (c *Client) handleWelcomeMessage(data []byte) {
	var payload SystemMessagePayload
	err := json.Unmarshal(data, &payload)
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
}
