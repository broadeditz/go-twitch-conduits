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
