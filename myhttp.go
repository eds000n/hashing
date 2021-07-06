package main

import (
	"flag"
	"fmt"
	"strings"
	"sync"

	"github.com/eds000n/hashing/http_client"
)

const ConcurrencyLimit = 10

// started at 8:22, finishing at 10:29
func main() {
	concurrency := flag.Int("parallel", ConcurrencyLimit, "maximum number of parallel requests")

	flag.Parse()

	urls := flag.Args()
	urls = preprocessURL(urls)

	limiter := make(chan string, *concurrency)
	var wg = &sync.WaitGroup{}
	wg.Add(len(urls))

	for _, url := range urls {
		go func(wg *sync.WaitGroup, limiter chan string, url string) {
			limiter <- "call"

			s, err := http_client.DoRequest(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%s %x\n", url, s)

			wg.Done()
			<-limiter
		}(wg, limiter, url)
	}
	wg.Wait()
}

// check if there's the http protocol prefixing the string
func preprocessURL(urls []string) []string {
	var fixedURLs []string
	for _, url := range urls {
		if !strings.HasPrefix(url, "http://") {
			fixedURLs = append(fixedURLs, "http://"+url)
		} else {
			fixedURLs = append(fixedURLs, url)
		}
	}
	return fixedURLs
}
