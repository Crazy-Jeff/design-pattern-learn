/*
过滤器模式(Filter Pattern)或标准模式（Criteria Pattern）是一种设计模式，这种模式允许开发人员使用不同的标准来过滤一组对象，通过逻辑运算以解耦的方式把它们连接起来。这种类型的设计模式属于结构型模式，它结合多个标准来获得单一标准。
	这种模式允许开发人员使用不同的标准来过滤一组对象，通过逻辑运算以解耦的方式把它们连接起来。这种类型的设计模式属于结构型模式，它结合多个标准来获得单一标准。
	source: https://www.runoob.com/design-pattern/filter-pattern.html
*/

package fitter

// Person 人
type Person interface {
	GetName() string
	GetAge() int
	Young() bool
}

// Male 男性
type Male struct {
	Name string
	Age  int
}

// Young 判断是否年轻
func (m *Male) Young() bool {
	return m.Age < 25
}

// GetName 获取名字
func (m *Male) GetName() string {
	return m.Name
}

// GetAge 获取年龄
func (m *Male) GetAge() int {
	return m.Age
}

// Female 女性
type Female struct {
	Name string
	Age  int
}

// Young 判断是否年轻
func (m *Female) Young() bool {
	return m.Age < 25
}

// GetName 获取名字
func (m *Female) GetName() string {
	return m.Name
}

// GetAge 获取年龄
func (m *Female) GetAge() int {
	return m.Age
}

// Youngster 过滤青少年
func Youngster(p ...Person) []Person {
	youngsters := make([]Person, 0, len(p))
	for i := range p {
		if p[i].Young() {
			youngsters = append(youngsters, p[i])
		}
	}
	return youngsters
}
