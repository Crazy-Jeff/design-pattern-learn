/*
责任链模式(Chain of Responsibility Pattern):
	意图：避免请求发送者与接收者耦合在一起，让多个对象都有可能接收请求，将这些对象连接成一条链，并且沿着这条链传递请求，直到有对象处理它为止。
	主要解决：职责链上的处理者负责处理请求，客户只需要将请求发送到职责链上即可，无须关心请求的处理细节和请求的传递，所以职责链将请求的发送者和请求的处理者解耦了。
	何时使用：在处理消息的时候以过滤很多道。
	如何解决：拦截的类都实现统一接口。
	关键代码：Handler 里面聚合它自己，在 HandlerRequest 里判断是否合适，如果没达到条件则向下传递，向谁传递之前 set 进去。
	应用实例： 1、红楼梦中的"击鼓传花"。 2、JS 中的事件冒泡。 3、JAVA WEB 中 Apache Tomcat 对 Encoding 的处理，Struts2 的拦截器，jsp servlet 的 Filter。
	优点： 1、降低耦合度。它将请求的发送者和接收者解耦。 2、简化了对象。使得对象不需要知道链的结构。 3、增强给对象指派职责的灵活性。通过改变链内的成员或者调动它们的次序，允许动态地新增或者删除责任。 4、增加新的请求处理类很方便。
	缺点： 1、不能保证请求一定被接收。 2、系统性能将受到一定影响，而且在进行代码调试时不太方便，可能会造成循环调用。 3、可能不容易观察运行时的特征，有碍于除错。
	使用场景： 1、有多个对象可以处理同一个请求，具体哪个对象处理该请求由运行时刻自动确定。 2、在不明确指定接收者的情况下，向多个对象中的一个提交一个请求。 3、可动态指定一组对象处理请求。
	注意事项：在 JAVA WEB 中遇到很多应用。
	source: https://www.runoob.com/design-pattern/chain-of-responsibility-pattern.html
*/

package cof

import "fmt"

// Country ...
type Country interface {
	Judge(string)
	SetNext(Country)
	Next() Country
}

// China ...
type China struct {
	nationality string
	next        Country
}

// Judge 国籍判断
func (c *China) Judge(nationality string) {
	if c.nationality == nationality {
		fmt.Println("这个人是中国人")
		return
	}
	if c.Next() != nil {
		c.Next().Judge(nationality)
	}
}

// SetNext 设置下个处理节点
func (c *China) SetNext(next Country) {
	c.next = next
}

// Next 获取下个处理节点
func (c *China) Next() Country {
	return c.next
}

// NewChinaHandler ...
func NewChinaHandler() *China {
	return &China{nationality: "CN"}
}

// Japan ...
type Japan struct {
	nationality string
	next        Country
}

// Judge 国籍判断
func (j *Japan) Judge(nationality string) {
	if j.nationality == nationality {
		fmt.Println("这个人是日本人")
		return
	}
	if j.Next() != nil {
		j.Next().Judge(nationality)
	}
}

// SetNext 设置下个处理节点
func (j *Japan) SetNext(next Country) {
	j.next = next
}

// Next 获取下个处理节点
func (j *Japan) Next() Country {
	return j.next
}

// NewJapanHandler ...
func NewJapanHandler() *Japan {
	return &Japan{nationality: "JP"}
}

// Korea ...
type Korea struct {
	nationality string
	next        Country
}

// Judge 国籍判断
func (k *Korea) Judge(nationality string) {
	if k.nationality == nationality {
		fmt.Println("这个人是韩国人")
		return
	}
	if k.Next() != nil {
		k.Next().Judge(nationality)
	}
}

// SetNext 设置下个处理节点
func (k *Korea) SetNext(next Country) {
	k.next = next
}

// Next 获取下个处理节点
func (k *Korea) Next() Country {
	return k.next
}

// NewKoreaHandler ...
func NewKoreaHandler() *Korea {
	return &Korea{nationality: "KR"}
}

// CountryHandler ...
type CountryHandler struct {
	chain Country
}

// Judge 国籍判断
func (ch *CountryHandler) Judge(nationality string) {
	ch.chain.Judge(nationality)
}

// NewCountryHandler 新建一个处理器
func NewCountryHandler() *CountryHandler {
	jp := NewJapanHandler()
	jp.SetNext(NewKoreaHandler())
	cn := NewChinaHandler()
	cn.SetNext(jp)
	return &CountryHandler{chain: cn}
}