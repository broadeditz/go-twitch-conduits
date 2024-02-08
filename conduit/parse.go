package conduit

import (
	"encoding/json"
	"errors"
	"io"
)

var (
	errUnexpectedToken = errors.New("unexpected token")
	errValueNotFound   = errors.New("value not found")
)

func ParseChannelMessage(data []byte) (ChannelMessage, error) {
	var msg ChannelMessage
	err := json.Unmarshal(data, &msg)
	return msg, err
}

// ParseMessageType parses the type from an event, without having to decode the entire message
func ParseMessageType(body io.Reader) (EventType, error) {
	d := json.NewDecoder(body)

	err := expectToken(d, json.Delim('{'))
	if err != nil {
		return "", err
	}

	// first we match the subscription element
	err = findValue(d, "subscription")
	if err != nil {
		if !errors.Is(err, errValueNotFound) {
			return "", err
		}
		return EventTypeNull, nil
	}

	// make sure the subscription element is an object
	err = expectToken(d, json.Delim('{'))
	if err != nil {
		return "", err
	}

	// then we match the type element
	err = findValue(d, "type")
	if err != nil {
		if !errors.Is(err, errValueNotFound) {
			return "", err
		}
		return EventTypeNull, nil
	}

	// get the type value
	t, err := d.Token()
	if err != nil {
		return "", err
	}

	eventType := ParseEventType(t.(string))

	return eventType, nil
}

func findValue(d *json.Decoder, key string) error {
	for d.More() {
		t, err := d.Token()
		if err != nil {
			return err
		}

		if t != key {
			err = skipValue(d)
			if err != nil {
				return err
			}

			continue
		}

		return nil
	}

	return errValueNotFound
}

func skipValue(d *json.Decoder) error {
	depth := 0

	for {
		t, err := d.Token()
		if err != nil {
			return err
		}

		switch t {
		case json.Delim('{'), json.Delim('['):
			depth++
		case json.Delim('}'), json.Delim(']'):
			depth--
		}

		if depth == 0 {
			return nil
		}
	}
}

func expectToken(d *json.Decoder, expected interface{}) error {
	t, err := d.Token()
	if err != nil {
		return err
	}

	if t != expected {
		return errUnexpectedToken
	}

	return nil
}
