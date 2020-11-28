package main

import (
	"github.com/ermos/hue"
	"log"
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

	err = bridge.Lights["11"].Set(hue.LightSetting{
		Alert: hue.AlertCycle,
		Xy: []float64{ 0.50, 0.80 },
	})
	if err != nil {
		log.Fatal(err)
	}
}
