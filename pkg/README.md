# pkg
This folder contains all go packages meant to be imported by other projects.

## Packages
- [conduit](./conduit): Package to manage a Twitch chat ingest using conduits
- [transport](./transport): Interface to manage websocket & webhook transports, along with code to parse incoming messages
- [webhook](./webhook): Logic for webhook conduits, implementing the `transport` interface
- [websocket](./websocket): Logic for websocket conduits, implementing the `transport` interface
