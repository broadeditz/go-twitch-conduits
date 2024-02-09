package helix

import (
	"io"
	"net/http"
)

// TwitchAPI is a struct used to interact with the Twitch helix API
type TwitchAPI struct {
	clientID   string
	oauthToken string
	userID     string
}

// NewTwitchAPI returns a new instance of the TwitchAPI configured with the given clientID, oauthToken, and userID
func NewTwitchAPI(clientID, oauthToken, userID string) *TwitchAPI {
	return &TwitchAPI{
		clientID:   clientID,
		oauthToken: oauthToken,
		userID:     userID,
	}
}

// UpdateOAuthToken updates the oauth token used by the TwitchAPI, used when the token expires
func (t *TwitchAPI) UpdateOAuthToken(oauthToken string) {
	t.oauthToken = oauthToken
}

// Do send a request to the given URL with the given method & body, adding authorization headers
func (t *TwitchAPI) Do(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Client-Id", t.clientID)
	req.Header.Set("Authorization", "Bearer "+t.oauthToken)
	req.Header.Set("Content-Type", "application/json")

	return http.DefaultClient.Do(req)
}
