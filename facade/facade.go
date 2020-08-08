/*
外观模式(Facade Pattern):
	意图：为子系统中的一组接口提供一个一致的界面，外观模式定义了一个高层接口，这个接口使得这一子系统更加容易使用。
	主要解决：降低访问复杂系统的内部子系统时的复杂度，简化客户端与之的接口。
	何时使用： 1、客户端不需要知道系统内部的复杂联系，整个系统只需提供一个"接待员"即可。 2、定义系统的入口。
	如何解决：客户端不与系统耦合，外观类与系统耦合。
	关键代码：在客户端和复杂系统之间再加一层，这一层将调用顺序、依赖关系等处理好。
	应用实例： 1、去医院看病，可能要去挂号、门诊、划价、取药，让患者或患者家属觉得很复杂，如果有提供接待人员，只让接待人员来处理，就很方便。 2、JAVA 的三层开发模式。
	优点： 1、减少系统相互依赖。 2、提高灵活性。 3、提高了安全性。
	缺点：不符合开闭原则，如果要改东西很麻烦，继承重写都不合适。
	使用场景： 1、为复杂的模块或子系统提供外界访问的模块。 2、子系统相对独立。 3、预防低水平人员带来的风险。
	注意事项：在层次化结构中，可以使用外观模式定义系统中每一层的入口。
	source: https://www.runoob.com/design-pattern/facade-pattern.html
*/

package facade

import "fmt"

// CarModel 模型
type CarModel struct {}

// SetModel 制作模型
func (cm *CarModel) SetModel() {
	fmt.Println("小王用塑料给您的新车打造了模型")
}

// CarTyre 轮胎
type CarTyre struct{}

// SetTyres 安装轮胎
func (ct *CarTyre) SetTyres() {
	fmt.Println("小李给您的新车装上了4个轮子,但是有一个忘记固定了")
}

// CarEngine 引擎
type CarEngine struct {}

// SetEngine 安装引擎
func (ce *CarEngine) SetEngine() {
	fmt.Println("小赵给您的新车装上了仓鼠跑步动能引擎")
}

// CarFactory 汽车工厂
type CarFactory struct {
	cm *CarModel
	ct *CarTyre
	ce *CarEngine
}

// NewCar 获取一台新车
func (cf *CarFactory) NewCar() {
	cf.cm.SetModel()
	cf.ct.SetTyres()
	cf.ce.SetEngine()
	fmt.Println("您得到一辆新车")
}	

// NewCarFactory 汽车工厂
func NewCarFactory() *CarFactory {
	return &CarFactory{
		cm: new(CarModel),
		ct: new(CarTyre),
		ce: new(CarEngine),
	}
}
