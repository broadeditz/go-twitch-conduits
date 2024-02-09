# helix

This package contains code to interface with the Twitch Helix API to manage subscribe to chat messages in conduits, as well as code to manage the chat messages themselves.  

## Usage

Call `helix.NewTwitchAPI` with a valid clientID, OAuth token and bot user ID.  
The twitch application of the clientID should have the `user:read:chat`, `user:bot` scopes for your bot user. And either the `channel:bot` scope, or moderator status for the channels you're trying to join.

## Missing features

- Conduits don't get deleted when all shards are disconnected.
- Sending messages to chat
- Handling any other chat message type than channel chat messages, such as whispers, follows, etc.
- Probably much more