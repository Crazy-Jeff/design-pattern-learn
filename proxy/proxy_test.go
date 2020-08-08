package proxy

import "testing"

func TestBuy(t *testing.T) {
	a := NewAgent()
	a.Buy()
}