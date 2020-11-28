package hue

import (
	"time"
)

func (l *Light) Alarm(ch chan error, duration time.Duration) error {
	currColor := l.State.Xy
	currState := l.State.On

	err := l.Toggle(true)
	if err != nil {
		return err
	}

	err = l.SetColorRGB(255, 0, 0)
	if err != nil {
		return err
	}

	d := false
	go func() {
		select {
		case <- time.After(duration):
			d = true
		case <- ch:
			d = true
		}
	}()

	for {
		if d {
			break
		}
		err = l.SetAlert(AlertCycle)
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}

	err = l.SetColorCIE(currColor[0], currColor[1])
	if err != nil {
		return err
	}

	return l.Toggle(currState)
}