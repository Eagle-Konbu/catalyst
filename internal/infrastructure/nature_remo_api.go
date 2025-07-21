package infrastructure

import (
	"fmt"
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
