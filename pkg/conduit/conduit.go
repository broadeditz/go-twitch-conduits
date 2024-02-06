package conduit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// CreateConduitResponse is the response the Twitch API returns when creating a conduit
type CreateConduitResponse struct {
	Data []struct {
		ID         string `json:"id"`
		ShardCount int    `json:"shard_count"`
	} `json:"data"`
}

// CreateConduit sends a request to the Twitch API to create a conduit with the given shard count
func (t *TwitchAPI) CreateConduit(shardCount int) (*CreateConduitResponse, error) {
	// POST to 'https://api.twitch.tv/helix/eventsub/conduits' with authorization & client ID headers, shard count in body
	requestData, body := map[string]interface{}{
		"shard_count": shardCount,
	}, new(bytes.Buffer)

	err := json.NewEncoder(body).Encode(requestData)
	if err != nil {
		return nil, err
	}

	res, err := t.Do(http.MethodPost, "https://api.twitch.tv/helix/eventsub/conduits", body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Create Conduit unexpected status code: %d", res.StatusCode))
	}

	var response CreateConduitResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type AssignConduitTransportRequest struct {
	ConduitID string           `json:"conduit_id"`
	Shards    []TransportShard `json:"shards"`
}

// TransportShard is a shard of a conduit
type TransportShard struct {
	ID        string `json:"id"`
	Status    string `json:"status,omitempty"`
	Transport struct {
		Method   TransportMethod `json:"method"`
		Callback string          `json:"callback"`
		Secret   string          `json:"secret"`
	}
}

type AssignConduitTransportResponse struct {
	Data []TransportShard `json:"data"`
}

// AssignConduitTransport sends a request to the Twitch API to assign a transport to shards of a conduit
func (t *TwitchAPI) AssignConduitTransport(request AssignConduitTransportRequest) (*AssignConduitTransportResponse, error) {
	// PATCH to 'https://api.twitch.tv/helix/eventsub/conduits/shards' with authorization & client ID headers, request body
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(request)
	if err != nil {
		return nil, err
	}

	res, err := t.Do(http.MethodPatch, "https://api.twitch.tv/helix/eventsub/conduits/shards", body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Assign Conduit Transport unexpected status code: %d", res.StatusCode))
	}

	var response AssignConduitTransportResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
