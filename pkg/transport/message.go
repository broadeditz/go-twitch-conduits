package transport

import "github.com/broadeditz/go-twitch-conduits/pkg/conduit"

type Subscription struct {
	ID        string            `json:"id"`
	Status    string            `json:"status"`
	Type      conduit.EventType `json:"type"`
	Version   string            `json:"version"`
	Condition struct {
		BroadcasterUserID string `json:"broadcaster_user_id"`
		UserID            string `json:"user_id"`
	} `json:"condition"`
	Transport struct {
		Method    conduit.TransportMethod `json:"method"`
		SessionID string                  `json:"session_id"`
	} `json:"transport"`
	CreatedAt string `json:"created_at"`
	Cost      int    `json:"cost"`
}

type ChannelMessage struct {
	Subscription Subscription        `json:"subscription"`
	Event        ChannelMessageEvent `json:"event"`
}

type ChannelMessageEvent struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	ChatterUserID        string `json:"chatter_user_id"`
	ChatterUserLogin     string `json:"chatter_user_login"`
	ChatterUserName      string `json:"chatter_user_name"`
	MessageID            string `json:"message_id"`
	Message              struct {
		Text      string `json:"text"`
		Fragments []struct {
			Type string `json:"type"`
			Text string `json:"text"`
			// TODO: figure out types
			Cheermote interface{} `json:"cheermote"`
			Emote     interface{} `json:"emote"`
			Mention   interface{} `json:"mention"`
		} `json:"fragments"`
	} `json:"message"`
	Color  string `json:"color"`
	Badges []struct {
		SetID string `json:"set_id"`
		ID    string `json:"id"`
		Info  string `json:"info"`
	} `json:"badges"`
	MessageType string `json:"message_type"`
	// TODO: figure out types
	Cheer                       interface{} `json:"cheer"`
	Reply                       interface{} `json:"reply"`
	ChannelPointsCustomRewardID interface{} `json:"channel_points_custom_reward_id"`
}
