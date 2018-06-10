// https://tour.golang.org/concurrency/10
package main

import (
	"fmt"
	"sync"
	"time"
)

type URLCache struct {
	visitedURL map[string]bool
	lock       sync.Mutex
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, urlCache *URLCache) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	(*urlCache).lock.Lock()
	_, visited := (*urlCache).visitedURL[url]
	if visited {
		fmt.Printf("already visited %v \n", url)
		(*urlCache).lock.Unlock()
		return
	} else {
		fmt.Printf("go to visit %v \n", url)
		(*urlCache).visitedURL[url] = true
		(*urlCache).lock.Unlock()
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, urlCache)
	}
	return
}

func main() {
	urlCache := URLCache{visitedURL: make(map[string]bool)}

	Crawl("https://golang.org/", 4, fetcher, &urlCache)
	time.Sleep(2 * time.Second)

	fmt.Println("urlCache", urlCache)

	fmt.Println("main end")
}
