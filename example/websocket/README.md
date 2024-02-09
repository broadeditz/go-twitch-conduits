# example/websocket

This example shows how to use the websocket client for twitch conduits.

## Usage

Requires 3 environment variables to be set:

- `CLIENT_ID`: The twitch client id for your application, which has the `user:read:chat`, `user:bot` for your bot user. And either `channel:bot` scope, or moderator status for the channels you're trying to join.
- `OAUTH_TOKEN`: An OAuth app token for the application, following the Client credential grant flow: https://dev.twitch.tv/docs/authentication/getting-tokens-oauth/#client-credentials-grant-flow
- `USER_ID`: The twitch user ID of the bot

Will console log all messages received in the bot's own channel