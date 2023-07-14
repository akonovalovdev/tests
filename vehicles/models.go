package vehicles

import (
	"errors"
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
