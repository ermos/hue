package main

import (
	"fmt"
	"github.com/ermos/hue"
	"log"
)

func main() {
	bridge := hue.Conn("192.168.1.2", hue.BridgeOptions{
		SaveToken: true,
		SaveLocation: "./",
		Debug: hue.DebugAll,
	})

	err := bridge.Fetch.Bridge()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bridge.Config.Name)
}
