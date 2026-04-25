package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Func func(string) ([]byte, error)

type result struct {
	url   string
	value []byte
}

type Memo struct {
	f     Func
	cache map[string]result
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func (memo *Memo) Get(url string) (*result, error) {
	res, ok := memo.cache[url]

	if !ok {
		byteRes, err := memo.f(url)
		if err != nil {
			return nil, err
		}

		myResult := result{
			url:   url,
			value: byteRes,
		}

		memo.cache[url] = myResult

		return &myResult, nil
	}

	return &res, nil
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
	incomingUrls := os.Args[1:]
	m := New(httpGetBody)

	for _, url := range incomingUrls {
		now := time.Now()
		_, err := m.Get(url)

		if err != nil {
			fmt.Printf("err in body parsing of url %s %s\n", url, err)
		}

		fmt.Printf("%s took %s\n", url, time.Since(now))
	}
}
