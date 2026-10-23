package main

import (
	"fmt"
	"sync"
)

type Post struct {
	views int
	mu    sync.Mutex
}

func (p *Post) init(wg *sync.WaitGroup) {
	defer wg.Done()
	p.mu.Lock()
	p.views += 1
	p.mu.Unlock()

}

func main() {
	var wg sync.WaitGroup
	myPost := Post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.init(&wg)

	}
	wg.Wait()
	fmt.Println(myPost.views)
}
