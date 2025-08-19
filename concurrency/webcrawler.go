// Q9: How can we parallelize a web crawler safely with concurrency?
package main

import (
	"fmt"
	"sync"
)

// FakeFetcher simulates a web fetcher
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

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

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{"https://golang.org/pkg/", "https://golang.org/cmd/"},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{"https://golang.org/", "https://golang.org/cmd/", "https://golang.org/pkg/fmt/"},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{"https://golang.org/", "https://golang.org/pkg/"},
	},
	"https://golang.org/cmd/": &fakeResult{
		"Commands",
		[]string{"https://golang.org/", "https://golang.org/pkg/"},
	},
}

var visited = struct {
	mu sync.Mutex
	m  map[string]bool
}{m: make(map[string]bool)}

func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	visited.mu.Lock()
	if visited.m[url] {
		visited.mu.Unlock()
		return
	}
	visited.m[url] = true
	visited.mu.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go func(link string) {
			defer wg.Done()
			Crawl(link, depth-1, fetcher)
		}(u)
	}
	wg.Wait()
}

func main() {
	Crawl("https://golang.org/", 3, fetcher)
}
