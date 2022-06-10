package main

import (
	"os"
	"strconv"
	"sync"
	"testing"
	"time"
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

	start := time.Now()

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
	t.Log(time.Now().Sub(start))
}

func TestGetRedirect(t *testing.T) {
	pool := generatePool()
	lc := loopCount()
	short := os.Getenv("SHORT")

	var wg sync.WaitGroup
	wg.Add(lc)

	start := time.Now()

	for i := 0; i < lc; i++ {
		pool <- func() {
			getRedirect(short)
			wg.Done()
		}
	}

	wg.Wait()
	t.Log(time.Now().Sub(start))
}

func TestGetRedirects(t *testing.T) {
	pool := generatePool()
	lc := loopCount()
	links := getLinks(lc)

	var wg sync.WaitGroup
	wg.Add(lc)

	start := time.Now()

	for i := 0; i < lc; i++ {
		pool <- func() {
			getRedirect(links.Data[i].Address)
			wg.Done()
		}
	}

	wg.Wait()
	t.Log(time.Now().Sub(start))
}

func TestGetLinks(t *testing.T) {
	limitInt := 10
	if limitStr := os.Getenv("LIMIT"); limitStr != "" {
		limitInt, _ = strconv.Atoi(limitStr)
	}

	getLinks(limitInt)
}
