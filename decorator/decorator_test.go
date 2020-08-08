package decorator

import (
	"fmt"
	"testing"
)

func TestWrapWork(t *testing.T) {
	f := WrapWork(Work)
	f("小赵", "搬运代码")
}

func TestWrapWorkWithExtra(t *testing.T) {
	f := WrapWorkWithExtra(Work, func(job string){
		fmt.Printf("为了更好的完成[%s],先到卫生间带薪拉屎...\n", job)
	})
	f("小赵", "搬运代码")
}

func TestWrapWorkWithMood(t *testing.T) {
	f := WrapWorkWithMood(Work, "开心地")
	f("小赵", "搬运代码")
}