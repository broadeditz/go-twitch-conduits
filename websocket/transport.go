package websocket

import (
	"github.com/gorilla/websocket"

	"github.com/broadeditz/go-twitch-conduits/helix"
)

// Client is a websocket implementation of the Transport interface
type Client struct {
	sessionID string

	conn *websocket.Conn

	ready     chan struct{}
	interrupt chan struct{}

	onChannelMessage func(message helix.ChannelMessage)
}

// NewClient returns a new instance of the websocket client
func NewClient() *Client {
	return &Client{
		ready:     make(chan struct{}),
		interrupt: make(chan struct{}),
	}
}

// Close attempts to gracefully close the websocket connection
func (c *Client) Close() {
	close(c.interrupt)
}

// GetTransportUpdate returns a helix.TransportUpdate for the websocket transport, used for updating conduits
func (c *Client) GetTransportUpdate() *helix.TransportUpdate {
	return &helix.TransportUpdate{
		Method:    helix.TransportMethodWebsocket,
		SessionID: c.sessionID,
	}
}

// Ready returns a channel that is closed when the websocket connection is ready
func (c *Client) Ready() chan struct{} {
	return c.ready
}

// OnChannelMessage sets a callback to be called when a channel message is received, is not executed in gorourines. It is the responsibility of the caller to handle concurrency.
func (c *Client) OnChannelMessage(f func(message helix.ChannelMessage)) {
	c.onChannelMessage = f
}
