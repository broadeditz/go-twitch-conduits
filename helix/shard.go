package helix

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// GetConduitsShardsResponse is the response the Twitch API returns when getting a list of shards for a conduit
type GetConduitsShardsResponse struct {
	Data       []GetConduitsShardsData `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

// GetConduitsShardsData is the data for a shard of a conduit
type GetConduitsShardsData struct {
	ID        string          `json:"id"`
	Status    string          `json:"status"`
	Transport TransportUpdate `json:"transport"`
}

// GetConduitShards gets a lists of all shards for a conduit.
func (t *TwitchAPI) GetConduitShards(conduitID string) (*GetConduitsShardsResponse, error) {
	// GET to 'https://api.twitch.tv/helix/eventsub/conduits/shards'
	res, err := t.Do(
		http.MethodGet,
		fmt.Sprintf("https://api.twitch.tv/helix/eventsub/conduits/shards?conduit_id=%v", conduitID),
		nil)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Get Conduits shards unexpected status code: %d, %+v", res.StatusCode, res.Status))
	}

	var response GetConduitsShardsResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// TODO: for v1.0.0 rename this to for UpdateConduitShards method instead
// AssignConduitTransportRequest is the request body for assigning a transport to shards of a conduit
type AssignConduitTransportRequest struct {
	ConduitID string           `json:"conduit_id"`
	Shards    []TransportShard `json:"shards"`
}

// TransportShard is a shard of a conduit
type TransportShard struct {
	ID        string          `json:"id"`
	Status    string          `json:"status,omitempty"`
	Transport TransportUpdate `json:"transport"`
}

// TransportUpdate contains the data of the transport for a conduit shard
// TODO: for v1.0.0 rename this to make more sense since it's used for more puposes than just updating a transport
type TransportUpdate struct {
	Method         TransportMethod `json:"method"`
	Callback       string          `json:"callback,omitempty"`
	ConduitID      string          `json:"conduit_id,omitempty"`
	SessionID      string          `json:"session_id,omitempty"`
	Secret         string          `json:"secret,omitempty"`
	ConnectedAt    string          `json:"connected_at,omitempty"`
	DisconnectedAt string          `json:"disconnected_at,omitempty"`
}

// GetConduitTransportRequest returns an AssignConduitTransportRequest for assigning a transport to a conduit using TwitchAPI.AssignConduitTransport
// TODO: for v1.0.0 rename
func (u *TransportUpdate) GetConduitTransportRequest(conduitID string, shardID string) *AssignConduitTransportRequest {
	return &AssignConduitTransportRequest{
		ConduitID: conduitID,
		Shards: []TransportShard{
			{
				ID:        shardID,
				Transport: *u,
			},
		},
	}
}

// AssignConduitTransportResponse is the response the Twitch API returns when assigning a transport a conduit
type AssignConduitTransportResponse struct {
	Data []TransportShard `json:"data"`
}

// UpdateConduitShards updates shard(s) for a conduit.
func (t *TwitchAPI) UpdateConduitShards(request *AssignConduitTransportRequest) (*AssignConduitTransportResponse, error) {
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

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusAccepted {
		return nil, errors.New(fmt.Sprintf("Assign Conduit Transport unexpected status code: %d, %+v", res.StatusCode, res.Status))
	}

	var response AssignConduitTransportResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// AssignConduitTransport sends a request to the Twitch API to assign a transport to shards of a conduit. Is an alias of UpdateConduitShards
// TODO: for v1.0.0 remove in favor of UpdateConduitShards
func (t *TwitchAPI) AssignConduitTransport(request *AssignConduitTransportRequest) (*AssignConduitTransportResponse, error) {
	return t.UpdateConduitShards(request)
}
