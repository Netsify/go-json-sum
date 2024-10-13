package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	numGoroutines := flag.Int("workers", 1, "number of goroutines to use")
	flag.Parse()

	file, err := os.Open("data.json")

	if err != nil {
		log.Fatalf("Error opening the file: %v", err)
	}

	defer file.Close()

	dataBytes, err := io.ReadAll(file)

	if err != nil {
		log.Fatalf("Error reading the file: %v", err)
	}

	var items []Item
	err = json.Unmarshal(dataBytes, &items)

	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	totalSum := sumElements(items, *numGoroutines)

	fmt.Printf("Total sum: %d\n", totalSum)
}
