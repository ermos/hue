package hue

import (
	"encoding/json"
	"errors"
	"io"
)

func isError(body io.Reader) bool {
	var e Error
	err := json.NewDecoder(body).Decode(&e)
	if err != nil || len(e) == 0 || e[0].Error.Type == 0 {
		return false
	}
	return true
}

func getError(body io.Reader) error {
	var e Error
	err := json.NewDecoder(body).Decode(&e)
	if err != nil {
		return err
	}
	return errors.New(e[0].Error.Description)
}