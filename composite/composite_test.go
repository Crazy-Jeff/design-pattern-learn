package composite

import (
	"fmt"
	"testing"
)

func TestShow(t *testing.T) {
	ms := make(map[string]*Menu)
	
	m1 := new(Menu)
	m1.AppendMenu(NewSubmenu("打开"))
	m1.AppendMenu(NewSubmenu("保存"))
	m1.AppendMenu(NewSubmenu("关闭"))
	ms["文件"] = m1

	m2 := new(Menu)
	m2.AppendMenu(NewSubmenu("复制"))
	m2.AppendMenu(NewSubmenu("剪切"))
	m2.AppendMenu(NewSubmenu("粘贴"))
	ms["编辑"] = m2


	for k, v := range ms {
		fmt.Printf("点击[%s]\n", k)
		v.Show()
	}

}