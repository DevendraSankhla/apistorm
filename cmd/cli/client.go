package main

import (
	"bytes"
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

func makePostCall(url string, wg *sync.WaitGroup, body string) {
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(body))
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
