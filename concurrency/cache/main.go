package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func httpGetBody(url string, c map[string][]byte) ([]byte, error) {
	//Reading from cache
	cachedRes, ok := c[url]
	if ok {
		return cachedRes, nil
	}

	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	red, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	//Caching it
	c[url] = red
	return red, nil
}

func main() {
	cacheSystem := make(map[string][]byte)

	incomingUrls := os.Args[1:]

	for _, url := range incomingUrls {
		now := time.Now()
		_, err := httpGetBody(url, cacheSystem)
		if err != nil {
			fmt.Printf("err in body parsing of url %s %s\n", url, err)
		}

		fmt.Printf("%s took %s\n", url, time.Since(now))
	}
}
