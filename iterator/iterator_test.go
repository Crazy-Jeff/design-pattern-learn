package iterator

import "testing"

func TestArray(t *testing.T) {
	arr := new(Array)
	arr.Append("abc", 1, 2, 3, true)
	for arr.Next() {
		t.Logf("index:%d, value:%v", arr.Index(), arr.Value())
	}
	arr.Append("OK", 3.1415)
	for arr.Next() {
		t.Logf("index:%d, value:%v", arr.Index(), arr.Value())
	}
}