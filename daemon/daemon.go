// daemon.go
package main

import (
	"fmt"
	"time"
)

func accumulate(s string, n int) int {
	sum := 0
	for ; n > 0; n-- {
		sum += n
		fmt.Printf("%s: %d\n", s, sum)
	}
	return sum
}

func main() {
	fmt.Printf("Total: %d\n", accumulate("A", 5))
	go accumulate("B", 4)
	time.Sleep(100 * time.Millisecond)
}
