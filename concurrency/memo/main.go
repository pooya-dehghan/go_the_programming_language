package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Func func(string) ([]byte, error)

type result struct {
	url   string
	value []byte
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(url string) ([]byte, error) {
	memo.mu.Lock()
	e := memo.cache[url]
	if e == nil {
		e = &entry{ready: make(chan struct{})}

		memo.cache[url] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(url)

		close(e.ready)
	} else {
		memo.mu.Unlock()
		<-e.ready
	}
	return e.res.value, e.res.err
}

func httpGetBody(url string) ([]byte, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	red, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return red, nil
}

func main() {
	incomingUrls := []string{"https://varzesh3.com", "https://isna.ir", "https://isna.ir", "https://saraf.app", "https://isna.ir", "https://isna.ir"}
	var wg sync.WaitGroup
	m := New(httpGetBody)

	for _, url := range incomingUrls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			now := time.Now()

			_, err := m.Get(u)

			if err != nil {
				fmt.Printf("err in body parsing of url %s %s\n", u, err)
			}

			fmt.Printf("%s took %s\n", u, time.Since(now))
		}(url)
	}

	wg.Wait()
}
