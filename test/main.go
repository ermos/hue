package main

import (
	"errors"
	"github.com/ermos/hue"
	"log"
	"time"
)

func main() {
	bridge := hue.Conn("192.168.1.2", hue.BridgeOptions{
		SaveToken: true,
		SaveLocation: "./",
		Debug: hue.DebugAll,
	})

	err := bridge.Fetch.Lights()
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan error)
	go func() {
		err = bridge.Lights["11"].Alarm(ch, 15 * time.Second)
		if err != nil {
			log.Fatal(err)
		}
	}()
	time.Sleep(5 * time.Second)
	ch <- errors.New("done")
	time.Sleep(5 * time.Second)
}
