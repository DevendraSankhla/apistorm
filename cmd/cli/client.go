package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func makeGetCall(url string, wg *sync.WaitGroup) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%s", body)
	defer wg.Done()
}
