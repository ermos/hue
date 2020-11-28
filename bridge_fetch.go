package hue

import (
	"bytes"
	"encoding/json"
	"github.com/ermos/hue/internal/logger"
)

func (b *BridgeFetch) Bridge() error {
	body, err := b.bridge.get("/config")
	if err != nil {
		return logger.Error(err)
	}

	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&b.bridge.Config)
	if err != nil {
		return logger.Error(err)
	}

	return nil
}
