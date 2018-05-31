package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64
var mutex sync.Mutex

func increment(s string) {
	for i := 0; i < 10; i++ {
		// without a lock to protect this critical section, each increment() execution thread will try to increment the counter
		// leading to a race and an incorrect final counter value
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		//mutex.Lock()
		//counter++
		//mutex.Unlock()

		// by using atomic variable the synchronization is taken care of by Go, the AddInt64 is
		// a thread safe operation
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go increment("Foo")
	go increment("Bar")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
