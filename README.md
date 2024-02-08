# go-twitch-conduits
Library to manage a Twitch chat ingest using conduits

## Packages
- [conduit](./conduit): Package to manage a Twitch chat ingest using conduits
- [webhook](./webhook): Logic for webhook conduits, implementing the `transport` interface
- [websocket](./websocket): Logic for websocket conduits, implementing the `transport` interface
- [oauth](./oauth): Oauth flow to let users give permission for the bot to join their channel