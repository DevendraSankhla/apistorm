package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

type config struct {
	totalCalls       int
	delayMillisecond int
}

func main() {
	var cfg config

	flag.IntVar(&cfg.totalCalls, "totalcalls", 0, "total calls to be made")
	flag.IntVar(&cfg.delayMillisecond, "delay", 0, "delay between each call in ms")

	flag.Parse()

	start := time.Now()
	var wg sync.WaitGroup
	for range cfg.totalCalls {
		wg.Add(1)
		go makeGetCall("http://localhost:4000/v1/healthcheck", &wg)
		time.Sleep(time.Duration(cfg.delayMillisecond) * time.Millisecond)
	}
	wg.Wait()
	elapsedTime := time.Since(start)
	fmt.Printf("Total calls: %d\nDelay: %d\nTotal elapsed time: %s", cfg.totalCalls, cfg.delayMillisecond, elapsedTime)
}
