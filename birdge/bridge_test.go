package bridge

import "testing"

func TestSayHello(t *testing.T) {
	r1 := ExecuteReincarnation("天蓬元帅", &Person{})
	r1.SayHello()

	r2 := ExecuteReincarnation("猪八戒", &Pig{})
	r2.SayHello()
}