package flyweight

import (
	"sync"
	"testing"
)

func TestGetPerson(t *testing.T) {
	pf := NewPersonFactory()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(2)
		go func(id int) {
			defer wg.Done()
			zs := pf.GetPerson("张三")
			t.Logf("id:%d, name:%s, uid:%s, pointer:%p, poolSize:%d", id, zs.Name, zs.ID, zs, len(pf.pool))
		}(i)
		go func(id int) {
			defer wg.Done()
			ls := pf.GetPerson("李四")
			t.Logf("id:%d, name:%s, uid:%s, pointer:%p, poolSize:%d", id, ls.Name, ls.ID, ls, len(pf.pool))
		}(i)
	}
	wg.Wait()
}

func BenchmarkNewPerson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Run("NewPerson", func(b *testing.B) {
			NewPerson("张三")
		})
	}
	pf := NewPersonFactory()
	for i := 0; i < b.N; i++ {
		b.Run("GetPerson", func(b *testing.B) {
			pf.GetPerson("张三")
		})
	}
}
