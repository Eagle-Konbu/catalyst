package usecase

import (
	"strconv"

	"github.com/Eagle-Konbu/catalyst/internal/infrastructure"
)

type AirconUsecase struct {
	id  string
	api *infrastructure.NatureRemoAPI
}

func NewAirconUsecase(api *infrastructure.NatureRemoAPI, id string) *AirconUsecase {
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
