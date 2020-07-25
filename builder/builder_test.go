package builder

import "testing"

func TestBurger_GetMeal(t *testing.T) {
	burger := NewBurger()
	cook := NewBurgerCook(burger)
	cook.Cook()
	product := burger.GetProduct()
	t.Log(product.GetPrice(), product.GetStep())
}