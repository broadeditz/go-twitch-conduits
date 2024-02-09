package websocket

import (
	"github.com/gorilla/websocket"

	"github.com/broadeditz/go-twitch-conduits/helix"
)

type Client struct {
	sessionID string

	conn *websocket.Conn

	ready     chan struct{}
	interrupt chan struct{}

	onChannelMessage func(message helix.ChannelMessage)
}

func NewClient() *Client {
	return &Client{
		ready:     make(chan struct{}),
		interrupt: make(chan struct{}),
	}
}

func (c *Client) Close() {
	close(c.interrupt)
}

func (c *Client) GetTransportUpdate() *helix.TransportUpdate {
	return &helix.TransportUpdate{
		Method:    helix.TransportMethodWebsocket,
		SessionID: c.sessionID,
	}
}

func (c *Client) Ready() chan struct{} {
	return c.ready
}

func (c *Client) OnChannelMessage(f func(message helix.ChannelMessage)) {
	c.onChannelMessage = f
}
