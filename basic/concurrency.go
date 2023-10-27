package main

import (
	"fmt"
	"time"
)

func say(s string, n int) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
}

func main() {
	go say("goroutine", 100)
	say("main", 10)
}
