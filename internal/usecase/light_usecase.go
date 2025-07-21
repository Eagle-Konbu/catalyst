package usecase

import "github.com/Eagle-Konbu/catalyst/internal/infrastructure"

type LightUsecase struct {
	API *infrastructure.NatureRemoAPI
}

func NewLightUsecase(token string) *LightUsecase {
	api := infrastructure.NewNatureRemoAPI(token)
	return &LightUsecase{API: api}
}

func (u *LightUsecase) SwitchLight(id string, button string) error {
	return u.API.SwitchLight(id, button)
}
