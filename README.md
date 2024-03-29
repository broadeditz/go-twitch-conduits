# go-twitch-conduits
Library to interact with Twitch chat messages using conduits in Go.  
Since the IRC will become more limited soon, this library is designed be a replacement ingest for chatbots and the likes.

## Packages
- [helix](./helix): Package to subscribe to chat messages through conduits using the Twitch Helix API
- [webhook](./webhook): Logic for webhook conduits, implementing the `helix.Transport` interface (TODO)
- [websocket](./websocket): Logic for websocket conduits, implementing the `helix.Transport` interface
- [oauth](./oauth): Oauth flow to let users give permission for the bot to join their channel (TODO)

## Usage

    go get github.com/broadeditz/go-twitch-conduits
  
Generally speaking, there are 3 things you need to get started:  
1. A Twitch application with the `user:read:chat`, `user:bot` scopes for your bot user. And either the `channel:bot` scope, or moderator status for the channels you're trying to join.
2. An OAuth app token for the application, following the [Client credential grant flow](https://dev.twitch.tv/docs/authentication/getting-tokens-oauth/#client-credentials-grant-flow).
3. The twitch user ID of the bot user
  
There are quick start examples in the `example` directory.

## v0.0.1 features

- [x] Subscribe to a chat (or more)
- [x] Websocket conduits
- [x] Transport interface for conduits & helix client
- [x] Small example of how to use the library

## Future roadmap

- [ ] Joining/leaving/deleting/updating an existing conduit
- [ ] Webhook conduits
- [ ] Oauth flow to let users give permission for the bot to join their channel
- [ ] Unit tests
- [ ] Single optional wrapper combining helix & transport into a simple to use package
- [ ] Send messages to chat
- [ ] v1.0.0: refactor helix to be more user-friendly