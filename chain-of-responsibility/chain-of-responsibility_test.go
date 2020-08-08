package cof

import "testing"

func TestJudge(t *testing.T) {
	ch := NewCountryHandler()
	ch.Judge("CN")
	ch.Judge("KR")
	ch.Judge("US") // 没有处理者进行处理
	ch.Judge("JP")
}