package hue

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func DiscoverAll() (dm []DiscoverModel, err error) {
	resp, err := http.Get("https://discovery.meethue.com")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &dm)
	return
}

func Discover() (dm DiscoverModel, err error) {
	dms, err := DiscoverAll()
	if err != nil {
		return
	}

	if len(dms) > 0 {
		return dms[0], nil
	}

	return
}