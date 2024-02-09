# go-twitch-conduits
Library to manage a Twitch chat ingest using conduits

## Packages
- [conduit](./conduit): Package to manage a Twitch chat ingest using conduits
- [webhook](./webhook): Logic for webhook conduits, implementing the `transport` interface (TODO)
- [websocket](./websocket): Logic for websocket conduits, implementing the `transport` interface
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