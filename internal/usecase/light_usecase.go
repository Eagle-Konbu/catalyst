package usecase

import "github.com/Eagle-Konbu/catalyst/internal/infrastructure"

type LightUsecase struct {
	id  string
	api *infrastructure.NatureRemoAPI
}

func NewLightUsecase(id, token string) *LightUsecase {
	api := infrastructure.NewNatureRemoAPI(token)
	return &LightUsecase{id: id, api: api}
}

func (u *LightUsecase) TurnOnLight() error {
	return u.api.SwitchLight(u.id, "on")
}

func (u *LightUsecase) TurnOffLight() error {
	return u.api.SwitchLight(u.id, "off")
}
