package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pinger(ping <-chan string, pong chan<- string) {
	for m := range ping {
		get(m)
		shot(pong, "pong")
	}
}

func ponger(ping chan<- string, pong <-chan string) {
	for m := range pong {
		get(m)
		shot(ping, "ping")
	}
}

func get(m string) {
	delay := time.Duration(rand.Intn(1000))
	fmt.Printf("got: %s, waiting: %dms\n", m, delay)
	time.Sleep(delay * time.Millisecond)
}

func shot(racket chan<- string, m string) {
	racket <- m
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go pinger(ping, pong)
	go ponger(ping, pong)
	ping <- "ping"

	for {
		// do nothing
	}
}
