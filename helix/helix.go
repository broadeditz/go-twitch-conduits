package helix

import (
	"io"
	"net/http"
)

type TwitchAPI struct {
	clientID   string
	oauthToken string
	userID     string
}

func NewTwitchAPI(clientID, oauthToken, userID string) *TwitchAPI {
	return &TwitchAPI{
		clientID:   clientID,
		oauthToken: oauthToken,
		userID:     userID,
	}
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
