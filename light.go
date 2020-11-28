package hue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ermos/hue/internal/logger"
	"github.com/lucasb-eyer/go-colorful"
)

type LightSetting struct {
	On        bool      `json:"on,omitempty"`
	Bri       int       `json:"bri,omitempty"`
	Hue       int       `json:"hue,omitempty"`
	Sat       int       `json:"sat,omitempty"`
	Effect    string    `json:"effect,omitempty"`
	Xy        []float64 `json:"xy,omitempty"`
	Ct        int       `json:"ct,omitempty"`
	Alert     string    `json:"alert,omitempty"`
}

const (
	AlertNone = "none"
	AlertCycle = "select"
	AlertCycle15Sec = "lselect"
	EffectNone = "none"
	EffectColorLoop = "colorloop"
)

// Allows you to set light color with RGB value
func (l *Light) SetColorRGB(r, g, b float64) error {
	if r > 255 || r < 0 {
		return logger.Error("red value is incorrect")
	}
	if g > 255 || g < 0 {
		return logger.Error("green value is incorrect")
	}
	if b > 255 || b < 0 {
		return logger.Error("blue value is incorrect")
	}

	c := colorful.Color{R: r, G: g, B: b}
	x, y, _ := c.Xyy()
	_, err := l.Bridge.put(
		fmt.Sprintf("/lights/%s/state", l.Key),
		bytes.NewBuffer(
			[]byte(
				fmt.Sprintf(`{"xy": [%f, %f]}`, x, y),
			),
		),
	)
	if err != nil {
		return logger.Error(err)
	}

	return nil
}

// Allows you to set light color with HEX value
func (l *Light) SetColorHEX(hex string) error {
	c, err := colorful.Hex(hex)
	if err != nil {
		return logger.Error(err)
	}

	x, y, _ := c.Xyy()

	_, err = l.Bridge.put(
		fmt.Sprintf("/lights/%s/state", l.Key),
		bytes.NewBuffer(
			[]byte(
				fmt.Sprintf(`{"xy": [%f, %f]}`, x, y),
			),
		),
	)
	if err != nil {
		return logger.Error(err)
	}

	return nil
}

// Allows to turn the light on and off, modify the hue and effects.
func (l *Light) Set(settings LightSetting) error {
	data, err := json.Marshal(settings)
	if err != nil {
		return logger.Error(err)
	}

	_, err = l.Bridge.put(
		fmt.Sprintf("/lights/%s/state", l.Key),
		bytes.NewReader(data),
	)
	if err != nil {
		return logger.Error(err)
	}

	return nil
}

// Used to rename lights. A light can have its name changed when in any state, including when it is unreachable or off.
func (l *Light) Rename(name string) error {
	if len(name) > 32 {
		return logger.Error("light name need to be between 0 and 32 characters")
	}

	_, err := l.Bridge.put(
		fmt.Sprintf("/lights/%s", l.Key),
		bytes.NewBuffer(
			[]byte(
				fmt.Sprintf(`{"name": "%s"}`, name),
				),
			),
		)
	if err != nil {
		return logger.Error(err)
	}

	return nil
}

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
	//{
	//    "7": {"name": "Hue Lamp 7"},
	//    "8": {"name": "Hue Lamp 8"},
	//    "lastscan": "2012-10-29T12:00:00"
	//}
}