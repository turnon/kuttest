package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const apiKey string = "2vZJAUhgmnDNqe0wGpNFm15VrTlxB_uK7307lFXp"

func main() {
}

func postLinks(params map[string]interface{}) {
	url := "http://localhost:3000/api/v2/links"

	bytes, _ := json.Marshal(params)
	str := string(bytes)

	payload := strings.NewReader(str)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("X-API-KEY", apiKey)
	req.Header.Add("Content-Type", "application/json")

	http.DefaultClient.Do(req)
}

type getLinksResp struct {
	Total int `json:"total"`
	Data  []struct {
		Address string `json:"address"`
	} `json:"data"`
}

func getLinks(limit int) *getLinksResp {
	url := "http://localhost:3000/api/v2/links?limit=" + strconv.Itoa(limit)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-API-KEY", apiKey)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var linksResp getLinksResp
	json.Unmarshal(body, &linksResp)

	return &linksResp
}

func getRedirect(short string) {
	url := "http://localhost:3000/" + short
	req, _ := http.NewRequest("GET", url, nil)
	http.DefaultClient.Do(req)
}

func randGenerator() func() uint64 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return func() uint64 {
		return r1.Uint64()
	}
}

func generatePool() chan<- func() {
	sizeInt := 1

	sizeStr := os.Getenv("POOL")
	if sizeStr != "" {
		sizeInt, _ = strconv.Atoi(sizeStr)
	}

	ch := make(chan func())
	for i := sizeInt; i > 0; i-- {
		go func() {
			for fn := range ch {
				fn()
			}
		}()
	}
	return ch
}

func loopCount() int {
	loopInt := 1000

	loopStr := os.Getenv("LOOP")
	if loopStr != "" {
		loopInt, _ = strconv.Atoi(loopStr)
	}

	return loopInt
}
