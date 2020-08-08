package facade

import "testing"

func TestNewCar(t *testing.T) {
	NewCarFactory().NewCar()
}