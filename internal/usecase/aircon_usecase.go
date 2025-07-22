package usecase

import (
	"strconv"

	"github.com/Eagle-Konbu/catalyst/internal/infrastructure"
)

type AirconUsecase struct {
	id  string
	api *infrastructure.NatureRemoAPI
}

func NewAirconUsecase(id, token string) *AirconUsecase {
	api := infrastructure.NewNatureRemoAPI(token)
	return &AirconUsecase{api: api, id: id}
}

func (u *AirconUsecase) SwitchAirconSettings(mode string, temp float64) error {
	var tempStr string
	if temp == float64(int64(temp)) {
		tempStr = strconv.FormatInt(int64(temp), 10)
	} else {
		tempStr = strconv.FormatFloat(temp, 'f', 1, 64)
	}
	return u.api.SwitchAirconSettings(u.id, mode, tempStr)
}

type AirconStatus struct {
	Mode        string
	Temperature float64
	TempUnit    string
}

func (u *AirconUsecase) GetAirconStatus() (*AirconStatus, error) {
	status, err := u.api.GetAirconStatus(u.id)
	if err != nil {
		return nil, err
	}
	return &AirconStatus{
		Mode:        status.Mode,
		Temperature: status.Temperature,
		TempUnit:    status.TempUnit,
	}, nil
}
