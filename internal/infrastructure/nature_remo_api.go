package infrastructure

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type NatureRemoAPI struct {
	Token string
}

func NewNatureRemoAPI(token string) *NatureRemoAPI {
	return &NatureRemoAPI{Token: token}
}

type PowerState string

const (
	PowerOn  PowerState = "on"
	PowerOff PowerState = "off"
)

// Reference: https://swagger.nature.global/#/default/post_1_appliances__applianceid__light
func (api *NatureRemoAPI) SwitchLight(id, button string) error {
	endpoint := fmt.Sprintf("https://api.nature.global/1/appliances/%s/light", id)
	data := url.Values{}
	data.Set("button", button)
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+api.Token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}
	return nil
}

// Reference: https://swagger.nature.global/#/default/post_1_appliances__applianceid__aircon_settings
func (api *NatureRemoAPI) SwitchAirconSettings(id, mode, temp string, power PowerState) error {
	endpoint := fmt.Sprintf("https://api.nature.global/1/appliances/%s/aircon_settings", id)
	data := url.Values{}
	if power == PowerOff {
		data.Set("button", "power-off")
	} else {
		if mode != "" {
			data.Set("operation_mode", mode)
		}
		if temp != "" {
			data.Set("temperature", temp)
		}
	}
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+api.Token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}
	return nil
}

// Appliance and AirconStatus structures for parsing response
type Appliance struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Settings *struct {
		Temp     string `json:"temp"`
		Mode     string `json:"mode"`
		TempUnit string `json:"temp_unit"`
	} `json:"settings"`
}

// Reference: https://swagger.nature.global/#/default/get_1_appliances
func (api *NatureRemoAPI) GetAppliances() ([]Appliance, error) {
	endpoint := "https://api.nature.global/1/appliances"
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+api.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var appliances []Appliance
	err = json.Unmarshal(body, &appliances)
	if err != nil {
		return nil, err
	}
	return appliances, nil
}
