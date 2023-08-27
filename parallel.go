package main

import (
	"fmt"
	"sync"
	"time"
)

const sliceSize = 100

func processElement(item int) int {
	// Simulate some time-consuming computation on each element
	time.Sleep(time.Millisecond * 10)
	return item * 2
}

func RunParallel() {
	// Create a slice with 10 million items
	slice := make([]int, sliceSize)
	for i := 0; i < sliceSize; i++ {
		slice[i] = i
	}

	numWorkers := 4 // Number of Goroutines (you can adjust this based on your system)
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Divide the work among Goroutines
	chunkSize := sliceSize / numWorkers

	for w := 0; w < numWorkers; w++ {
		startIndex := w * chunkSize
		endIndex := (w + 1) * chunkSize
		if w == numWorkers-1 {
			// The last Goroutine might have more elements if sliceSize is not divisible by numWorkers
			endIndex = sliceSize
		}

		go func(start, end int) {
			defer wg.Done()
			for i := start; i < end; i++ {
				slice[i] = processElement(slice[i])
			}
		}(startIndex, endIndex)
	}

	wg.Wait()

	fmt.Println("Processing is done.")
	// You can use the processed slice here.
}
