package api_calls

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"spotify/handlers/models"
)

type Device struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DevicesResponse struct {
	Devices []Device `json:"devices"`
}

func Get_Devices(client *http.Client, token string, trackInfo models.TrackInfo) (string, error) {
	url := "https://api.spotify.com/v1/me/player/devices"
	defDevice := "(no active device)"

	req, err := prepareDeviceRequest(url, token)
	if err != nil {
		return defDevice, err
	}

	resp, err := sendDeviceRequest(client, req)
	if err != nil {
		return defDevice, err
	}

	devicesResp, err := parseDeviceResponse(resp.Body)
	if err != nil {
		return defDevice, err
	}

	return handleDeviceData(client, devicesResp, trackInfo.TrackID, token, defDevice)
}

func prepareDeviceRequest(url, token string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	setDevicesHeaders(req, token)
	return req, nil
}

func sendDeviceRequest(client *http.Client, req *http.Request) (*http.Response, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 400 {
		return nil, fmt.Errorf("received error status: %d", resp.StatusCode)
	}

	return resp, nil
}

func parseDeviceResponse(body io.Reader) (DevicesResponse, error) {
	var devicesResp DevicesResponse
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&devicesResp)
	if err != nil {
		return devicesResp, err
	}
	return devicesResp, nil
}

func handleDeviceData(client *http.Client, devicesResp DevicesResponse, trackID, token, defDevice string) (string, error) {
	if len(devicesResp.Devices) > 0 {
		primaryDevice := devicesResp.Devices[0].ID
		name := devicesResp.Devices[0].Name
		err := playTrack(client, trackID, primaryDevice, token)
		return name, err
	}
	return defDevice, nil
}

func setDevicesHeaders(req *http.Request, token string) {
	req.Header.Set("Authorization", "Bearer "+token)
}
