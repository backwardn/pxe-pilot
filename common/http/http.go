package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"dev.splitted-desktop.com/horizon/pxe-pilot/logger"
)

// HTTP do an HTTP call
func HTTP(method string, baseURL string, path string, data interface{}, responseHolder interface{}) (int, error) {

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(data)

	url := fmt.Sprintf("%s%s", baseURL, path)

	logger.Info(" -> Sending %s request on %s", method, url)

	req, err := http.NewRequest(method, url, b)
	if err != nil {
		return -1, err
	}

	req.Close = true
	req.Header.Set("Content-Type", "application/json")

	var transport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	client := http.Client{
		Transport: transport,
		Timeout:   time.Duration(5 * time.Second),
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if responseHolder != nil {
		json.NewDecoder(resp.Body).Decode(responseHolder)
	}

	logger.Info(" -> Response code %d", resp.StatusCode)

	return resp.StatusCode, nil
}