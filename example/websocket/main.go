package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/broadeditz/go-twitch-conduits/helix"
	"github.com/broadeditz/go-twitch-conduits/websocket"
)

var (
	clientID   = os.Getenv("CLIENT_ID")
	oauthToken = os.Getenv("OAUTH_TOKEN")
	userID     = os.Getenv("USER_ID")
)

func main() {
	api := helix.NewTwitchAPI(clientID, oauthToken, userID)
	transport := helix.Transport(websocket.NewClient())

	res, err := api.CreateConduit(1)
	if err != nil {
		panic(err)
	}

	// register the message handler
	transport.OnChannelMessage(handleMessage)

	errChan := make(chan error)
	go func() {
		// Initialize the websocket connection
		errChan <- transport.Init()
	}()

	// wait for the transport to be ready
	<-transport.Ready()

	// Get the conduit update request for the shard
	req := transport.GetTransportUpdate().GetConduitTransportRequest(res.Data[0].ID, "0")

	// Register the shard in the conduit
	_, err = api.AssignConduitTransport(req)
	if err != nil {
		panic(err)
	}

	// Subscribe to the channel.chat.message event
	_, err = api.EventSubscribeChannelMessage(res.Data[0].ID, userID, userID)
	if err != nil {
		panic(err)
	}

	// listen for system signal to exit
	close := make(chan os.Signal)
	signal.Notify(close, os.Interrupt, os.Kill)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				panic(err)
			}

			fmt.Printf("Graceful shut down\n")
			return

		case <-close:
			transport.Close()
		}
	}
}

func handleMessage(msg helix.ChannelMessage) {
	fmt.Printf("%+v\n", msg)
}
