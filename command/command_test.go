package command

import "testing"

func TestExec(t *testing.T) {
	cs := NewCommands()
	cs.Exec("eating")
	cs.Exec("running")
	cs.Exec("shopping")
	cs.Exec("sleeping")
}