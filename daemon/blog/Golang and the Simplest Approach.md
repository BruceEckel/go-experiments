## It might not scale...

Go’s default mode around concurrency appears to be “the simplest approach,” which means:

1. Every goroutine is a daemon thread. Ignoring task completion, cancellation, and error handling is certainly the simplest approach. Tasks are terminated when the program ends, so it doesn’t matter if terminating a task puts the program in an undefined state. If you need completion, cancellation and/or error handling, you add it yourself with more complicated code.

2. Every goroutine only uses communicating sequential processes (CSP). This prevents data races with the least amount of cognitive overhead. It works great as long as you’re not moving so much data that it impacts your program. Again, for the majority of Go programs this is a fine default, and again, you use other approaches by adding more complicated code.

The statement "every goroutine is a daemon thread" might sound odd at first, because every concurrency system has a way to start a task, and that task will keep running until it either ends on its own or is terminated by an external agent. What makes something a daemon is that it "runs in the background." In particular, daemons don't have a connection with the task/process/thread that started them. In Unix, daemons are typically started by a process that then exits, leaving the daemon with no parent process to control it. It just keeps running in the background.

Let's look at how Go does things:
``` go
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
```
Following the normal model for functions, `accumulate()` takes arguments and returns a result. In the first line of `main()`, we call `accumulate()` and display that result. But in the second line, we use the `go` statement, and *this statement doesn't return anything*. I get no "handle" back from starting a goroutine, which means I have no connection with that goroutine which would allow me to cancel it or *to receive its return value*.
## The Easy Concurrency Model

The defaults (1 & 2 at the beginning of this post) are a great fit for a certain set of programs. And if Go’s goal is to enable programmers to “program concurrently without knowing about concurrency,” then I must admit they’ve done an admirable job. But it seems to me that Go has been pitched (or at least, absorbed by the Go community) as a general purpose language suitable for all applications.

Concurrency seems like the area where I have seen the most incidences of the Dunning-Kruger effect (you learn a little and think you know a lot). There are so many different strategies/niches in the concurrency world that it is extremely easy to learn one of those and immediately think that you understand concurrency. A big giveaway here is if someone declares concurrency to be “easy.” Understand that this person has reached a happy place where they’ve gotten something working after being told that concurrency is a big, scary thing. Telling them “actually, concurrency is still big and scary” will not be gratefully received—they want to stay in that happy place and you’re rocking the boat.

I wonder what happens when all these Go programmers who have been existing in this happy concurrency world encounter its limits. The logical thing to do would be to understand Go’s limits and choose a different technology when it no longer fits. But that requires knowing enough about concurrency so you can see and understand those limits.

If you’re invested in Go, you’ll want to keep using Go. You incrementally learn more and add more complexity to your program as needed. Like all incremental complexity, this seems perfectly rational in the moment—why not add a little bit of complexity rather than making fundamental technology changes? Lots of projects totter along like this, regularly incrementing their complexity, and never reach the “aha” moment when they realize they’ve pushed that technology far beyond its boundaries. Will the company see the necessity of rewriting the project so that it can continue to expand, or will they continue to complexify the current project one feature at the time?

From a cultural standpoint, Go’s approach might be excellent. Get everyone started with a basic, relatively-foolproof model. Let them get experience with concurrency before moving on to more complicated issues. This is generally how we teach, to keep from intimidating and overwhelming the students. And remember that a main reason for the success of C++ is that you can just program in C, giving C programmers a gentler path into the language than if the entire world suddenly changes. This was important in an age when C was the only language someone might have learned (that person often having started in assembly language).

My experience of concurrency is that it is in a different universe than programming languages. It is a leaky abstraction and you can easily encounter operating system or even hardware details. Concurrency is a collection of strategies, and these are often dramatically different from each other.

Learning the (maybe) easiest approach (daemons + CSP) is certainly a reasonable way to onboard new concurrency programmers, but I am curious to see what happens after time, when projects mature enough to hit the boundaries of that approach.