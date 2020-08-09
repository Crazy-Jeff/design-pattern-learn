/*
解释器模式（Interpreter Pattern）提供了评估语言的语法或表达式的方式，它属于行为型模式。这种模式实现了一个表达式接口，该接口解释一个特定的上下文。这种模式被用在 SQL 解析、符号处理引擎等。
	意图：给定一个语言，定义它的文法表示，并定义一个解释器，这个解释器使用该标识来解释语言中的句子。
	主要解决：对于一些固定文法构建一个解释句子的解释器。
	何时使用：如果一种特定类型的问题发生的频率足够高，那么可能就值得将该问题的各个实例表述为一个简单语言中的句子。这样就可以构建一个解释器，该解释器通过解释这些句子来解决该问题。
	如何解决：构建语法树，定义终结符与非终结符。
	关键代码：构建环境类，包含解释器之外的一些全局信息，一般是 HashMap。
	应用实例：编译器、运算表达式计算。
	优点： 1、可扩展性比较好，灵活。 2、增加了新的解释表达式的方式。 3、易于实现简单文法。
	缺点： 1、可利用场景比较少。 2、对于复杂的文法比较难维护。 3、解释器模式会引起类膨胀。 4、解释器模式采用递归调用方法。
	使用场景： 1、可以将一个需要解释执行的语言中的句子表示为一个抽象语法树。 2、一些重复出现的问题可以用一种简单的语言来进行表达。 3、一个简单语法需要解释的场景。
	注意事项：可利用场景比较少，JAVA 中如果碰到可以用 expression4J 代替。
	source: https://www.runoob.com/design-pattern/interpreter-pattern.html
*/

package interpreter

import (
	"errors"
	"strings"
)

// Expression ...
type Expression interface {
	Integer(Expression) int64
}

// Variable ...
type Variable struct {
	value string
}

// Integer ...
func (i *Variable) Integer(e Expression) int64 {
	var sum int64
	for j := range []rune(i.value) {
		sum += int64([]rune(i.value)[j])
	}
	return sum
}

// Add ,,,
type Add struct {
	result int64
}

// Integer ...
func (a *Add) Integer(e Expression) int64 {
	if e != nil {
		a.result += e.Integer(e)
	}
	return a.result
}

// Sub ,,,
type Sub struct {
	first  bool
	result int64
}

// Integer ...
func (s *Sub) Integer(e Expression) int64 {
	if e != nil {
		if s.first {
			s.result = e.Integer(e)
			s.first = false
		} else {
			s.result -= e.Integer(e)
		}
	}
	return s.result
}

// Evaluator ...
type Evaluator struct {
	placeholder int
	operate     Expression
	params      []Expression
}

// Exec ...
func (e *Evaluator) Exec(params ...string) (int64, error) {
	l := len(params)
	if l != e.placeholder {
		return 0, errors.New("参数个数不匹配")
	}

	for i := range e.params {
		if int64(rune('?')) == e.params[i].Integer(e.operate) {
			e.operate.Integer(&Variable{value: params[l-e.placeholder]})
			e.placeholder--
		} else {
			e.operate.Integer(e.params[i])
		}
	}

	return e.operate.Integer(nil), nil
}

// NewEvaluator ...
func NewEvaluator(src string) (*Evaluator, error) {
	cmds := strings.Split(src, " ")
	if len(cmds) < 3 {
		return nil, errors.New("最少需要3个参数 \"{操作符} 参数1 参数2 ...\"")
	}

	evaluator := new(Evaluator)
	evaluator.params = make([]Expression, len(cmds)-1)
	switch cmds[0] {
	case "+":
		evaluator.operate = new(Add)
	case "-":
		evaluator.operate = &Sub{first: true}
	default:
		return nil, errors.New("不支持的运算命令")
	}

	for i := range cmds[1:] {
		switch cmds[i+1] {
		case "?":
			evaluator.placeholder++
		}
		evaluator.params[i] = &Variable{value: cmds[i+1]}
	}
	return evaluator, nil
}
