package helix

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// EventSubscribeRequest is the request body for subscribing to an eventSub event
type EventSubscribeRequest struct {
	Type      EventType               `json:"type"`
	Version   string                  `json:"version"`
	Condition EventSubscribeCondition `json:"condition"`
	Transport TransportUpdate         `json:"transport"`
}

// EventSubscribeCondition contains the user data for the event subscription
type EventSubscribeCondition struct {
	BroadcasterUserID string `json:"broadcaster_user_id"`
	ModeratorUserID   string `json:"moderator_user_id,omitempty"`
	UserID            string `json:"user_id,omitempty"`
}

// EventSubscribeResponse is the response the Twitch API returns when subscribing to an eventSub event
type EventSubscribeResponse struct {
	Data []struct {
		ID        string                  `json:"id"`
		Status    string                  `json:"status"`
		Type      EventType               `json:"type"`
		Version   string                  `json:"version"`
		Condition EventSubscribeCondition `json:"condition"`
		CreatedAt string                  `json:"created_at"`
		Transport TransportUpdate         `json:"transport"`
		Cost      int                     `json:"cost"`
	}
	Total        int `json:"total"`
	MaxTotalCost int `json:"max_total_cost"`
	TotalCost    int `json:"total_cost"`
}

// EventSubscribeChannelMessage subscribes to chat events in the given channel, as the given user. This is more or less the equivalent of JOIN in IRC.
func (t *TwitchAPI) EventSubscribeChannelMessage(conduitID, channelID, userID string) (*EventSubscribeResponse, error) {
	request := GetChannelMessageSubscribeRequest(conduitID, channelID, userID)
	return t.EventSubscribe(request)
}

// GetChannelMessageSubscribeRequest returns an EventSubscribeRequest for subscribing to chat events in the given channel, as the given user
func GetChannelMessageSubscribeRequest(conduitID, channelID, userID string) *EventSubscribeRequest {
	return &EventSubscribeRequest{
		Type:    EventTypeChannelMessage,
		Version: "1",
		Condition: EventSubscribeCondition{
			BroadcasterUserID: channelID,
			UserID:            userID,
		},
		Transport: TransportUpdate{
			Method:    TransportMethodConduit,
			ConduitID: conduitID,
		},
	}
}

// EventSubscribe sends a request to the Twitch API to subscribe to an eventSub event
func (t *TwitchAPI) EventSubscribe(request *EventSubscribeRequest) (*EventSubscribeResponse, error) {
	// POST to 'https://api.twitch.tv/helix/eventsub/subscriptions' with authorization & client ID headers, request body
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(request)
	if err != nil {
		return nil, err
	}

	res, err := t.Do(http.MethodPost, "https://api.twitch.tv/helix/eventsub/subscriptions", body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusAccepted {
		return nil, errors.New(fmt.Sprintf("Event Subscribe unexpected status code: %d, %+v", res.StatusCode, res.Status))
	}

	var response EventSubscribeResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
