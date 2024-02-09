# go-twitch-conduits
Library interact with Twitch chat messages using conduits in Go.  
Since IRC will become more limited soon, this library is designed be a replacement ingest for chatbots and the likes.

## Packages
- [helix](./helix): Package to manage subscribe to chat messages in conduits using the Twitch Helix API
- [webhook](./webhook): Logic for webhook conduits, implementing the `Transport` interface (TODO)
- [websocket](./websocket): Logic for websocket conduits, implementing the `Transport` interface
- [oauth](./oauth): Oauth flow to let users give permission for the bot to join their channel (TODO)

## Usage

Generally speaking, there are 3 things you need to get started:  
1. A Twitch application with the `user:read:chat`, `user:bot` and either the `channel:bot` scope, or moderator status for the channels you're trying to join.
2. An OAuth app token for the application, following the [Client credential grant flow](https://dev.twitch.tv/docs/authentication/getting-tokens-oauth/#client-credentials-grant-flow)
3. The twitch user ID of the bot
  
There are quick start examples in the `example` directory.

## v0.0.1 requirements

- [x] Subscribe to a chat
- [x] Websocket conduits
- [ ] Well thought out spec for transport interface & helix client
- [x] Small example of how to use the library
- [ ] Documentation/comments

## Future roadmap

- [ ] Webhook conduits
- [ ] Oauth flow to let users give permission for the bot to join their channel
- [ ] Unit tests
- [ ] Single optional wrapper combining helix & transport into a simple to use package