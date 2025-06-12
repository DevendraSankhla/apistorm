package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

type config struct {
	url              string
	totalCalls       int
	delayMillisecond int
	method           string
	body             string
}

func main() {
	var cfg config

	flag.StringVar(&cfg.method, "method", "", "REST api method")
	flag.StringVar(&cfg.url, "url", "", "Api url to call")
	flag.StringVar(&cfg.body, "postbody", "{}", "Post reqest body")
	flag.IntVar(&cfg.totalCalls, "totalcalls", 0, "Total calls to be made")
	flag.IntVar(&cfg.delayMillisecond, "delay", 0, "Delay between each call in ms")

	flag.Parse()

	start := time.Now()
	var wg sync.WaitGroup
	for range cfg.totalCalls {
		wg.Add(1)
		if cfg.method == "GET" {
			go makeGetCall(cfg.url, &wg)
		} else if cfg.method == "POST" {
			go makePostCall(cfg.url, &wg, cfg.body)
		} else {
			panic("need api method")
		}
		time.Sleep(time.Duration(cfg.delayMillisecond) * time.Millisecond)
	}
	wg.Wait()
	elapsedTime := time.Since(start)
	fmt.Printf("Total calls: %d\nDelay: %d\nTotal elapsed time: %s", cfg.totalCalls, cfg.delayMillisecond, elapsedTime)
}
