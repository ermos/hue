package hue


import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ermos/hue/internal/logger"
	"io/ioutil"
	"net/http"
)

func (b *BridgeFetch) Lights() error {
	resp, err := http.Get(fmt.Sprintf("http://%s/api/%s/lights", b.Bridge.IPAddr, b.Bridge.Token))
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

	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&b.Bridge.Lights)
	if err != nil {
		return logger.Error(err)
	}

	return nil
}
