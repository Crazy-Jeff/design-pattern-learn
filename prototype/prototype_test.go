package prototype

import "testing"

func TestNewPerson(t *testing.T) {
	p1 := NewPerson("老王")
	p2 := p1.Clone()
	t.Logf("p1名:%s, p2名:%s", p1.Name(), p2.Name())
	p2.SetName("老张")
	t.Logf("p1名:%s, p2名:%s", p1.Name(), p2.Name())
	p1.SetName("老李")
	t.Logf("p1名:%s, p2名:%s", p1.Name(), p2.Name())
}