package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type result struct {
	url        string
	statusCode int
	duration   time.Duration
}

func worker(wg *sync.WaitGroup, jobs chan string, res chan result) {
	defer wg.Done()
	for url := range jobs {
		start := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("%s err: %s\n", url, err)
			continue
		}
		duration := time.Since(start)
		res <- result{url, resp.StatusCode, duration}
		resp.Body.Close()
	}
}

func main() {
	url := []string{
		"http://baidu.com",
		"http://shixiaocaia.fun",
	}

	numG := 5
	var wg sync.WaitGroup
	wg.Add(numG)

	jobs := make(chan string, numG)
	result := make(chan result, numG)

	for i := 1; i <= 5; i++ {
		go worker(&wg, jobs, result)
	}

	for i := 0; i < len(url); i++ {
		jobs <- url[i]
	}
	close(jobs)

	wg.Wait()
	close(result)

	for ans := range result {
		fmt.Printf("url:%s, statusCode:%d, duration:%v\n", ans.url, ans.statusCode, ans.duration)
	}
}
