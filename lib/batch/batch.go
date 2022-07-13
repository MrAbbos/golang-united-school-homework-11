package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {

	var mx sync.Mutex
	var wg sync.WaitGroup
	var i int64 = 0

	ch := make(chan struct{}, pool)
	for ; i < n; i++ {
		wg.Add(1)
		ch <- struct{}{}
		go func(i int64) {
			el := getOne(i)
			mx.Lock()
			res = append(res, el)
			mx.Unlock()
			<-ch
			wg.Done()
		}(i)
	}
	wg.Wait()
	return res
}
