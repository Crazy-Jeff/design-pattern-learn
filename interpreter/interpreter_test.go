package interpreter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecAdd(t *testing.T) {
	evaluator, err := NewEvaluator("+ ? ABC ? 好")
	if err != nil {
		t.Error(err)
		return
	}
	result, err := evaluator.Exec("c", "F")
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, result, int64(rune('c')+rune('A')+rune('B')+rune('C')+rune('F')+rune('好')))
}


func TestExecSub(t *testing.T) {
	evaluator, err := NewEvaluator("- ? ABC ? 好")
	if err != nil {
		t.Error(err)
		return
	}
	result, err := evaluator.Exec("c", "F")
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, result, int64(rune('c')-rune('A')-rune('B')-rune('C')-rune('F')-rune('好')))
}