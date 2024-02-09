package websocket

import (
	"github.com/gorilla/websocket"

	"github.com/broadeditz/go-twitch-conduits/conduit"
)

type Client struct {
	sessionID string

	conn *websocket.Conn

	ready     chan struct{}
	interrupt chan struct{}

	onChannelMessage func(message conduit.ChannelMessage)
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

func (c *Client) GetTransportUpdate() *conduit.TransportUpdate {
	return &conduit.TransportUpdate{
		Method:    conduit.TransportMethodWebsocket,
		SessionID: c.sessionID,
	}
}

func (c *Client) Ready() chan struct{} {
	return c.ready
}

func (c *Client) OnChannelMessage(f func(message conduit.ChannelMessage)) {
	c.onChannelMessage = f
}
