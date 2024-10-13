package main

import "sync"

func sumElements(items []Item, numGoroutines int) int {
	var wg sync.WaitGroup
	resultChan := make(chan int, numGoroutines)

	chunkSize := (len(items) + numGoroutines - 1) / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize

		if start >= len(items) {
			break
		}

		wg.Add(1)

		go func(start int) {
			defer wg.Done()
			sum := 0
			end := start + chunkSize

			if end > len(items) {
				end = len(items)
			}

			for _, item := range items[start:end] {
				sum += item.A + item.B
			}

			resultChan <- sum
		}(start)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	totalSum := 0

	for partialSum := range resultChan {
		totalSum += partialSum
	}

	return totalSum
}
