package main

import (
	"fmt"
	"sync"
)

func calculatesquare(num int, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	square := num * num
	results <- square
}

func main() {
	numbers := []int{2, 10, 4, 8, 6}

	results := make(chan int, len(numbers))

	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go calculatesquare(num, results, &wg)

	}

	wg.Wait()

	close(results)
	fmt.Println("Square of numbers:")
	for square := range results {
		fmt.Println(square)
	}
}
