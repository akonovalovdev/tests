package vehicles

import (
	"encoding/json"
	"errors"
	"time"
)

var (
	PetrolError = errors.New("not enough fuel, visit a petrol station")
	GasError    = errors.New("not enough fuel, visit a gas station")
)

type TaxiDriver struct {
	Vehicle     Vehicle `json:"-"`
	ID          int     `json:"id"`
	OrdersCount int     `json:"orders"`
}

func (x *TaxiDriver) SetVehicle(isEvening bool) {
	if !isEvening {
		x.Vehicle = &Camry{
			FuelConsumption: 10,
			EngineLeft:      1000,
			IsPetrol:        true,
		}
	} else {
		x.Vehicle = &LandCruiser{
			FuelConsumption: 16,
			EngineLeft:      2000,
			IsPetrol:        false,
		}
	}
}

func (x *TaxiDriver) Drive() error {
	if err := x.Vehicle.ConsumeFuel(); err != nil {
		return err
	}

	x.OrdersCount++
	return nil
}

type ReportData struct {
	TaxiDriver
	Date time.Time `json:"date"`
}

func (x *TaxiDriver) SendDailyReport() ([]byte, error) {
	data := ReportData{
		TaxiDriver: *x,
		Date:       time.Now(),
	}

	msg, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	x.OrdersCount = 0
	return msg, nil
}

type Vehicle interface {
	ConsumeFuel() error
}

type Camry struct {
	FuelConsumption float32
	EngineLeft      float32
	IsPetrol        bool
}

func (x *Camry) ConsumeFuel() error {
	if x.FuelConsumption > x.EngineLeft {
		return PetrolError
	}

	x.EngineLeft -= x.FuelConsumption
	return nil
}

type LandCruiser struct {
	FuelConsumption float32
	EngineLeft      float32
	IsPetrol        bool
}

func (x *LandCruiser) ConsumeFuel() error {
	if x.FuelConsumption > x.EngineLeft {
		return GasError
	}

	x.EngineLeft -= x.FuelConsumption
	return nil
}
