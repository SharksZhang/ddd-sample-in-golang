package service

import (
    "github.com/agiledragon/ddd-sample-in-golang/domain/model"
    "sync"
)

type CargoService struct {
    repo model.CargoRepo
    provider model.CargoProvider
}

var cs = &CargoService{}
var once sync.Once
func GetCargoServiceInstance() *CargoService {
    once.Do(func() {
        cs.repo = model.GetCargoRepo()
        cs.provider = model.GetCargoProvider()
    })
    return cs
}

func (t *CargoService) Create(cargoId string, afterDays uint) {
    cargo := model.CargoFactory{}.Create(cargoId, afterDays)
    t.repo.Add(cargo)
    t.provider.Confirm(cargo)
}

func (t *CargoService) Delay(cargoId string, days uint) {
    cargo := t.repo.Get(cargoId)
    if cargo != nil {
        cargo.Delay(days)
        t.repo.Update(cargo)
        t.provider.Confirm(cargo)
    }
}

func (t *CargoService) GetAfterDays(cargoId string) uint {
    cargo := t.repo.Get(cargoId)
    if cargo != nil {
        return cargo.GetAfterDays()
    }
    panic("invalid cargoId")
}