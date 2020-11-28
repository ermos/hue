package hue

import (
	"bytes"
	"fmt"
	"github.com/ermos/hue/internal/logger"
	"github.com/lucasb-eyer/go-colorful"
)

const (
	AlertNone = "none"
	AlertCycle = "select"
	AlertCycle15Sec = "lselect"
	EffectNone = "none"
	EffectColorLoop = "colorloop"
)

// Toggle light to ON/OFF
func (l *Light) Toggle(on bool) error {
	_, err := l.Bridge.put(
		fmt.Sprintf("/lights/%s/state", l.Key),
		bytes.NewBuffer(
			[]byte(
				fmt.Sprintf(`{"on": %t}`, on),
			),
		),
	)
	if err != nil {
		return logger.Error(err)
	}

	l.State.On = on

	return nil
}

// Set brightness to the light
func (l *Light) SetBrightness(bri uint8) error {
	_, err := l.Bridge.put(
		fmt.Sprintf("/lights/%s/state", l.Key),
		bytes.NewBuffer(
			[]byte(
				fmt.Sprintf(`{"bri": %d}`, bri),
			),
		),
	)
	if err != nil {
		return logger.Error(err)
	}

	l.State.Bri = bri

	return nil
}

// Set Hue to the light
func (l *Light) SetHue(hue uint16) error {
	_, err := l.Bridge.put(
		fmt.Sprintf("/lights/%s/state", l.Key),
		bytes.NewBuffer(
			[]byte(
				fmt.Sprintf(`{"hue": %d}`, hue),
			),
		),
	)
	if err != nil {
		return logger.Error(err)
	}

	l.State.Hue = hue

	return nil
}

// Set saturation of the light
func (l *Light) SetSaturation(sat uint8) error {
	_, err := l.Bridge.put(
		fmt.Sprintf("/lights/%s/state", l.Key),
		bytes.NewBuffer(
			[]byte(
				fmt.Sprintf(`{"sat": %d}`, sat),
			),
		),
	)
	if err != nil {
		return logger.Error(err)
	}

	l.State.Sat = sat

	return nil
}

// Set effect to light, actually only "colorloop" exist
func (l *Light) SetEffect(effect string) error {
	_, err := l.Bridge.put(
		fmt.Sprintf("/lights/%s/state", l.Key),
		bytes.NewBuffer(
			[]byte(
				fmt.Sprintf(`{"effect": "%s"}`, effect),
			),
		),
	)
	if err != nil {
		return logger.Error(err)
	}

	l.State.Effect = effect

	return nil
}

// Set alert effect to light
func (l *Light) SetAlert(alert string) error {
	_, err := l.Bridge.put(
		fmt.Sprintf("/lights/%s/state", l.Key),
		bytes.NewBuffer(
			[]byte(
				fmt.Sprintf(`{"alert": "%s"}`, alert),
			),
		),
	)
	if err != nil {
		return logger.Error(err)
	}

	l.State.Alert = alert

	return nil
}

// Allows you to set color temperature
func (l *Light) SetColorTemperature(ct uint16) error {
	_, err := l.Bridge.put(
		fmt.Sprintf("/lights/%s/state", l.Key),
		bytes.NewBuffer(
			[]byte(
				fmt.Sprintf(`{"ct": %d}`, ct),
			),
		),
	)
	if err != nil {
		return logger.Error(err)
	}

	l.State.Ct = ct

	return nil
}

// Allows you to set light color with CIE value
func (l *Light) SetColorCIE(x, y float64) error {
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

	l.State.Xy = []float64{ x, y }

	return nil
}

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

	l.State.Xy = []float64{ x, y }

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

	l.State.Xy = []float64{ x, y }

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

	l.Name = name

	return nil
}