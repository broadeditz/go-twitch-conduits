# go-twitch-conduits
Library to subscribe to Twitch chat messages using conduits in Go
Since IRC will become more limited soon, this library is designed be a replacement ingest for chatbots and the likes.

## Packages
- [helix](./helix): Package to manage subscribe to chat messages in conduits using the Twitch Helix API
- [webhook](./webhook): Logic for webhook conduits, implementing the `Transport` interface (TODO)
- [websocket](./websocket): Logic for websocket conduits, implementing the `Transport` interface
- [oauth](./oauth): Oauth flow to let users give permission for the bot to join their channel (TODO)

## v0.0.1 requirements

- [x] Subscribe to a chat
- [x] Websocket conduits
- [ ] Well thought out spec for transport interface & helix client
- [ ] Small example of how to use the library
- [ ] Documentation/comments

## Future roadmap

- [ ] Webhook conduits
- [ ] Oauth flow to let users give permission for the bot to join their channel
- [ ] Unit tests
- [ ] Single optional wrapper combining helix & transport into a simple to use package