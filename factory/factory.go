package factory

import (
	"fmt"
)

// Restaurant 声明一个饭店类型
type Restaurant interface {
	Cook()
}

// GetFood 获取一份食物
func GetFood(name string) Restaurant {
	switch name {
	case "面条":
		return new(Noodles)
	case "米饭":
		return new(Rice)
	default:
		return new(Tea)
	}
}

// Noodles 菜品:面条
type Noodles struct {
}

// Cook 制作一份面条
func (n *Noodles) Cook() {
	fmt.Println("得到一份面条")
}

// Rice 菜品:米饭
type Rice struct {
}

// Cook 制作一份米饭
func (r *Rice) Cook() {
	fmt.Println("得到一份米饭")
}

// Tea 菜品:茶水
type Tea struct {
}

// Cook 获取一杯茶水
func (t *Tea) Cook() {
	fmt.Println("不知道您要吃什么,先喝杯茶吧")
}

// SimpleRestaurant 定义简单的餐厅
type SimpleRestaurant struct {
}

// NewSimpleRestaurant 获取一家饭店
func NewSimpleRestaurant() *SimpleRestaurant {
	return new(SimpleRestaurant)
}

// GetRice 获取米饭
func (s *SimpleRestaurant) GetRice() Restaurant {
	return new(Rice)
}

// GetNoodles 获取面条
func (s *SimpleRestaurant) GetNoodles() Restaurant {
	return new(Noodles)
}

// GetTea 获取茶水
func (s *SimpleRestaurant) GetTea() Restaurant {
	return new(Tea)
}