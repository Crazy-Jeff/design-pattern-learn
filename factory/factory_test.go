package factory

import (
	"testing"
)

func TestCook(t *testing.T) {
	GetFood("").Cook()
	GetFood("米饭").Cook()
	GetFood("面条").Cook()
}

func TestNewSimpleRestaurant(t *testing.T) {
	r := NewSimpleRestaurant()
	r.GetNoodles().Cook()
	r.GetRice().Cook()
	r.GetTea().Cook()
}