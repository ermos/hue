package hue

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func (b *Bridge) get(url string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s/api/%s%s", b.IPAddr, b.Token, url))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if isError(bytes.NewBuffer(body)) {
		return nil, getError(bytes.NewBuffer(body))
	}

	return body, nil
}

func (b *Bridge) put(url string, body io.Reader) ([]byte, error) {
	return b.action(http.MethodPut, url, body)
}

func (b *Bridge) post(url string, body io.Reader) ([]byte, error) {
	return b.action(http.MethodPost, url, body)
}

func (b *Bridge) action(method string, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(
		method,
		fmt.Sprintf("http://%s/api/%s%s", b.IPAddr, b.Token, url),
		body,
	)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	if isError(bytes.NewBuffer(buf.Bytes())) {
		return nil, getError(bytes.NewBuffer(buf.Bytes()))
	}

	return buf.Bytes(), nil
}