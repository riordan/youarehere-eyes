package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/schollz/wifiscan"
)

type HerePlatform struct {
	AppId   string `json:"app_id"`
	AppCode string `json:"app_code"`
}

type PositionResponse struct {
	Location struct {
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lng"`
		Accuracy  int     `json:"accuracy"`
	} `json:"location"`
}

func (platform *HerePlatform) position(payload []byte) (PositionResponse, error) {
	endpoint, _ := url.Parse("https://pos.api.here.com/positioning/v1/locate")
	queryParams := endpoint.Query()
	queryParams.Set("app_id", platform.AppId)
	queryParams.Set("app_code", platform.AppCode)
	endpoint.RawQuery = queryParams.Encode()
	response, err := http.Post(endpoint.String(), "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return PositionResponse{}, err
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var positionResponse PositionResponse
		json.Unmarshal(data, &positionResponse)
		return positionResponse, nil
	}
}

func scanWifiNetworks() []wifiscan.Wifi {
	wifis, err := wifiscan.Scan()
	if err != nil {
		log.Fatal(err)
	}
	return wifis
}

func main() {
	networks := scanWifiNetworks()
	for _, w := range networks {
		formatMac := strings.ReplaceAll(strings.ToUpper(w.SSID), ":", "-")
		fmt.Println(formatMac, w.RSSI)
	}

}
