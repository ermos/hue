package hue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ermos/hue/internal/logger"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	DebugNone = 0
	DebugInfo = 1
	DebugError = 2
	DebugAll = 3
)

type BridgeOptions struct {
	Token 			string
	SaveToken 		bool
	SaveLocation 	string
	Debug 			int
}

// Make new conn to the bridge
func Conn(IPAddr string, options BridgeOptions) *Bridge {
	b := &Bridge{
		IPAddr: IPAddr,
	}

	// Initialize Fetch
	b.Fetch = BridgeFetch{
		bridge: b,
	}

	// Initialize debug
	if options.Debug != 0 {
		logger.SetLevel(options.Debug)
	}

	// Check if token already exist
	if options.Token != "" {
		b.Token = options.Token
	}
	if options.SaveToken {
		if _, err := os.Stat(filepath.Join(options.SaveLocation, ".hue")); !os.IsNotExist(err) {
			tokenByte, err := ioutil.ReadFile(filepath.Join(options.SaveLocation, ".hue"))
			if err != nil {
				log.Fatal(err)
			}
			b.Token = string(tokenByte)
		}
	}

	// Don't have token ? Wait for button press
	if b.Token == "" {
		for {
			if err := b.auth(); err == nil {
				break
			}
			time.Sleep(2 * time.Second)
		}
	}

	// Save token
	if options.SaveToken {
		err := ioutil.WriteFile(filepath.Join(options.SaveLocation, ".hue"), []byte(b.Token), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	return b
}

func (b *Bridge) Rename(name string) error {
	if len(name) < 4 || len(name) > 16 {
		return logger.Error("bridge name need to be between 4 and 16 characters")
	}

	_, err := b.put(
		"/config",
		bytes.NewBuffer(
			[]byte(
				fmt.Sprintf(`{"name": "%s"}`, name),
			),
		),
	)
	if err != nil {
		return logger.Error(err)
	}

	b.Config.Name = name

	return nil
}

func (b *Bridge) auth() error {
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("http://%s/api", b.IPAddr),
		bytes.NewBuffer([]byte(fmt.Sprintf(`{"devicetype": "hue-%d"}`, time.Now().Unix()))),
		)
	if err != nil {
		return logger.Error(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return logger.Error(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return logger.Error(err)
	}

	if isError(bytes.NewBuffer(body)) {
		return logger.Error(getError(bytes.NewBuffer(body)))
	}

	var token tokenResponse
	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&token)
	if err != nil {
		return logger.Error(err)
	}

	b.Token = token[0].Success.Username

	return nil
}