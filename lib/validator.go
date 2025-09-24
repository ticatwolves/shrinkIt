package lib

import (
	"encoding/json"
	"errors"
	schema "shrinkIt/internal/schemas"
)

func RequestParser(body []byte, requestPayload *schema.RequestPayload) error {
	err := json.Unmarshal([]byte(body), &requestPayload)
	if err != nil {
		return err
	}
	if requestPayload.Url == "" {
		return errors.New("URL can not be empty")
	}
	return nil
}
