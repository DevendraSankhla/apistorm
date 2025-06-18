package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sync"
)

func makeAPICall(url string, method string, wg *sync.WaitGroup, body string) {
	req, err := http.NewRequest(method, url, bytes.NewBufferString(body))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%s", resBody)

	defer wg.Done()
	defer res.Body.Close()
}
