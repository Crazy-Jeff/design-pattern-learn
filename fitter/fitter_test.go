package fitter

import "testing"

func TestYoungster(t *testing.T) {
	youngsters := Youngster(&Male{Name: "老王", Age: 33}, &Female{Name: "小张", Age: 13}, &Male{Name: "小明", Age: 22}, &Male{Name: "小赵", Age: 24}, &Female{Name: "李梅梅", Age: 28}, &Female{Name: "小刘", Age: 7})
	if len(youngsters) == 0 {
		t.Log("没有青少年")
	} else {
		for i := range youngsters {
			t.Logf("年轻者: 名字:%s, 年龄:%d", youngsters[i].GetName(), youngsters[i].GetAge())
		}
	}
}