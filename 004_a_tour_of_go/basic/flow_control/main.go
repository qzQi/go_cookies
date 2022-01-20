package main

import (
	"fmt"
)

func main() {
	// for
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	// The init and post statements are optional.
	for sum < 1000 {
		sum += sum
	}

	// the while
	for sum < 2000 {
		sum += sum
	}

	// forever
	for {
		break
	}
	fmt.Println(sum)
}
