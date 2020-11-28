package hue

import (
	"bytes"
	"encoding/json"
	"github.com/ermos/hue/internal/logger"
)

func (b *BridgeFetch) Lights() error {
	body, err := b.Bridge.get("/lights")
	if err != nil {
		return logger.Error(err)
	}

	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&b.Bridge.Lights)
	if err != nil {
		return logger.Error(err)
	}

	for key, light := range b.Bridge.Lights {
		light.Key = key
		light.Bridge = b.Bridge
	}

	return nil
}
