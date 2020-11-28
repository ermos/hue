package hue

import (
	"bytes"
	"encoding/json"
	"github.com/ermos/hue/internal/logger"
)

func (b *BridgeFetch) Lights() error {
	body, err := b.bridge.get("/lights")
	if err != nil {
		return logger.Error(err)
	}

	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&b.bridge.Lights)
	if err != nil {
		return logger.Error(err)
	}

	for key, light := range b.bridge.Lights {
		light.Key = key
		light.Bridge = b.bridge
	}

	return nil
}
