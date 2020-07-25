/*
工厂模式(Factory Pattern):
	意图：定义一个创建对象的接口,并在子类中实现该接口的功能,可以根据需求使用不同的子类进行生产
	主要解决：主要解决接口选择的问题。
	何时使用：我们明确地计划不同条件下创建不同实例时。
	如何解决：让其子类实现工厂接口，返回的也是一个抽象的产品。
	关键代码：创建过程在其子类执行。
	应用实例：用户想要购买一份食物,那么可以通知饭店,饭店根据用户的需求让不同的厨师进行烹饪,用户只要提供需要的食物类型,然后等待拿走食物即可,不需要关心食物是餐厅里哪位厨师制作,怎样制作的.
	优点： 1、一个调用者想创建一个对象，只要知道其名称就可以了。 2、扩展性高，如果想增加一个产品，只要扩展一个工厂类就可以。 3、屏蔽产品的具体实现，调用者只关心产品的接口。
	缺点：每次增加一个产品时，都需要增加一个具体类和对象实现工厂，使得系统中类的个数成倍增加，在一定程度上增加了系统的复杂度，同时也增加了系统具体类的依赖。这并不是什么好事。
	使用场景： 1、日志记录器：记录可能记录到本地硬盘、系统事件、远程服务器等，用户可以选择记录日志到什么地方。 2、数据库访问，当用户不知道最后系统采用哪一类数据库，以及数据库可能有变化时。 3、设计一个连接服务器的框架，需要三个协议，"POP3"、"IMAP"、"HTTP"，可以把这三个作为产品类，共同实现一个接口。
	注意事项：作为一种创建类模式，在任何需要生成复杂对象的地方，都可以使用工厂方法模式。有一点需要注意的地方就是复杂对象适合使用工厂模式，而简单对象，特别是只需要通过 new 就可以完成创建的对象，无需使用工厂模式。如果使用工厂模式，就需要引入一个工厂类，会增加系统的复杂度。
	source: https://www.runoob.com/design-pattern/factory-pattern.html

抽象工厂模式(Abstract Factory Pattern):
	意图：提供一个创建一系列相关或相互依赖对象的接口，而无需指定它们具体的类。
	主要解决：主要解决接口选择的问题。
	何时使用：系统的产品有多于一个的产品族，而系统只消费其中某一族的产品。
	如何解决：在一个产品族里面，定义多个产品。
	关键代码：在一个工厂里聚合多个同类产品。
	应用实例：工作了，为了参加一些聚会，肯定有两套或多套衣服吧，比如说有商务装（成套，一系列具体产品）、时尚装（成套，一系列具体产品），甚至对于一个家庭来说，可能有商务女装、商务男装、时尚女装、时尚男装，这些也都是成套的，即一系列具体产品。假设一种情况（现实中是不存在的，要不然，没法进入共产主义了，但有利于说明抽象工厂模式），在您的家中，某一个衣柜（具体工厂）只能存放某一种这样的衣服（成套，一系列具体产品），每次拿这种成套的衣服时也自然要从这个衣柜中取出了。用 OOP 的思想去理解，所有的衣柜（具体工厂）都是衣柜类的（抽象工厂）某一个，而每一件成套的衣服又包括具体的上衣（某一具体产品），裤子（某一具体产品），这些具体的上衣其实也都是上衣（抽象产品），具体的裤子也都是裤子（另一个抽象产品）。
	优点：当一个产品族中的多个对象被设计成一起工作时，它能保证客户端始终只使用同一个产品族中的对象。
	缺点：产品族扩展非常困难，要增加一个系列的某一产品，既要在抽象的 Creator 里加代码，又要在具体的里面加代码。
	使用场景： 1、QQ 换皮肤，一整套一起换。 2、生成不同操作系统的程序。
	注意事项：产品族难扩展，产品等级易扩展。
	source: https://www.runoob.com/design-pattern/abstract-factory-pattern.html
*/

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
