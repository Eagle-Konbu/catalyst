package infrastructure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type NatureRemoAPI struct {
	Token string
}

func NewNatureRemoAPI(token string) *NatureRemoAPI {
	return &NatureRemoAPI{Token: token}
}

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
func (api *NatureRemoAPI) SwitchAirconSettings(id, mode, temp string) error {
	endpoint := fmt.Sprintf("https://api.nature.global/1/appliances/%s/aircon_settings", id)
	data := url.Values{}
	if mode != "" {
		data.Set("operation_mode", mode)
	}
	if temp != "" {
		data.Set("temperature", temp)
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

type AirconStatus struct {
	Mode        string
	Temperature float64
	TempUnit    string
}

// GetAppliances fetches the list of appliances from Nature Remo API
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
	body, err := ioutil.ReadAll(resp.Body)
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

// GetAirconStatus fetches the mode and temperature for the given acId
func (api *NatureRemoAPI) GetAirconStatus(acId string) (*AirconStatus, error) {
	appliances, err := api.GetAppliances()
	if err != nil {
		return nil, err
	}
	for _, a := range appliances {
		if a.ID == acId && a.Type == "AC" && a.Settings != nil {
			temp, _ := strconv.ParseFloat(a.Settings.Temp, 64)
			return &AirconStatus{
				Mode:        a.Settings.Mode,
				Temperature: temp,
				TempUnit:    a.Settings.TempUnit,
			}, nil
		}
	}
	return nil, fmt.Errorf("AC with id %s not found", acId)
}
