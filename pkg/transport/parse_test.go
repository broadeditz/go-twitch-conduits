package transport

import (
	"bytes"
	"io"
	"testing"

	"github.com/broadeditz/go-twitch-conduits/pkg/conduit"
)

func TestParseMessageType(t *testing.T) {
	type args struct {
		body io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    conduit.EventType
		wantErr bool
	}{
		{
			name: "TestParseChannelMessage",
			args: args{
				bytes.NewReader([]byte(`
{
  "subscription": {
    "id": "0b7f3361-672b-4d39-b307-dd5b576c9b27",
    "status": "enabled",
    "type": "channel.chat.message",
    "version": "1",
    "condition": {
      "broadcaster_user_id": "1971641",
      "user_id": "2914196"
    },
    "transport": {
      "method": "websocket",
      "session_id": "AgoQHR3s6Mb4T8GFB1l3DlPfiRIGY2VsbC1h"
    },
    "created_at": "2023-11-06T18:11:47.492253549Z",
    "cost": 0
  },
  "event": {
    "broadcaster_user_id": "1971641",
    "broadcaster_user_login": "streamer",
    "broadcaster_user_name": "streamer",
    "chatter_user_id": "4145994",
    "chatter_user_login": "viewer32",
    "chatter_user_name": "viewer32",
    "message_id": "cc106a89-1814-919d-454c-f4f2f970aae7",
    "message": {
      "text": "Hi chat",
      "fragments": [
        {
          "type": "text",
          "text": "Hi chat",
          "cheermote": null,
          "emote": null,
          "mention": null
        }
      ]
    },
    "color": "#00FF7F",
    "badges": [
      {
        "set_id": "moderator",
        "id": "1",
        "info": ""
      },
      {
        "set_id": "subscriber",
        "id": "12",
        "info": "16"
      },
      {
        "set_id": "sub-gifter",
        "id": "1",
        "info": ""
      }
    ],
    "message_type": "text",
    "cheer": null,
    "reply": null,
    "channel_points_custom_reward_id": null
  }
}`),
				),
			},
			want:    conduit.EventTypeChannelMessage,
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMessageType(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMessageType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseMessageType() got = %v, want %v", got, tt.want)
			}
		})
	}
}
