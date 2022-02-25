package main

import (
	"fmt"
	"time"
)

func main() {
	list := []int{3, 5, 2, 11, 4, 7, 1, 8, 0, 13, 9, 6}
	fmt.Printf("before: %v -> sorted:%v\n", list, sleepSort(list))
}

// Exercise: Implement sleep sort
// For every value 'n' in values spin a goroutine that will
// sleep 'n' milliseconds and then send the value on a channel
func sleepSort(values []int) []int {
	out := make([]int, 0, len(values))
	ch := make(chan int)

	for _, val := range values {
		go func(v int) {
			// With n times 20 milliseconds, it manages to sort nicely
			time.Sleep(time.Duration(v) * 20 * time.Millisecond)
			ch <- v
		}(val)
	}

	for i := 0; i < len(values); i++ {
		val := <-ch
		out = append(out, val)
	}

	return out
}
