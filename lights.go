package hue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ermos/hue/internal/logger"
)

// The bridge will open the network for 40s and search all available devices, send multiple time append 40s.
// For show new devices, you need to use ShowNewLights.
func (b *Bridge) SearchNewLights() error {
	_, err := b.post(
		fmt.Sprintf("/lights"),
		nil,
	)
	if err != nil {
		return logger.Error(err)
	}

	return nil
}

// Gets a list of lights that were discovered the last time a search for new lights was performed.
// The list of new lights is always deleted when a new search is started.
// If you miss one light, you can find it in group "0"
func (b *Bridge) ShowNewLights() (list []AvailableLight, err error) {
	body, err := b.get("/lights/new")
	if err != nil {
		return list, logger.Error(err)
	}

	var parse map[string]interface{}
	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&parse)
	if err != nil {
		return list, logger.Error(err)
	}
	delete(parse, "lastscan")

	for id, name := range parse {
		list = append(list, AvailableLight{
			ID: id,
			Name: name.(string),
		})
	}

	if len(list) != 0 {
		err = b.Fetch.Lights()
		if err != nil {
			return list, err
		}
	}

	return list, nil
}