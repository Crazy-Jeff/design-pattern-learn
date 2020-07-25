package singleton

import (
	"sync"
	"testing"
)

func TestGetPerson(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetPerson().IncrAge()
		}()
	}
	wg.Wait()
	p := GetPerson()
	t.Logf("name: %s, birthday: %s, now: %s, age: %d", p.GetName(), p.GetBirthday().Format("2006-01-02"), p.GetNow().Format("2006-01-02"), p.GetAge())
}