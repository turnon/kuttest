package main

import (
	"strconv"
	"sync"
	"testing"
)

func TestRandGenerator(t *testing.T) {
	randGen := randGenerator()

	for i := 0; i < 10; i++ {
		t.Log(randGen())
	}
}

func TestPostLinks(t *testing.T) {
	randGen := randGenerator()
	pool := generatePool()
	lc := loopCount()

	var wg sync.WaitGroup
	wg.Add(lc)

	for i := 0; i < lc; i++ {
		pool <- func() {
			num := strconv.FormatUint(randGen(), 10)
			target := "https://baijiahao.baidu.com/s?id=" + num
			postLinks(map[string]interface{}{
				"target": target,
			})
			wg.Done()
		}
	}

	wg.Wait()
}

func TestGetRedirect(t *testing.T) {
	randGen := randGenerator()
	pool := generatePool()
	lc := loopCount()

	for i := 0; i < lc; i++ {
		pool <- func() {
			num := strconv.FormatUint(randGen(), 10)
			getRedirect(num)
		}
	}
}
