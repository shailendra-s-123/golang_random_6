package main

import (
	"fmt"
	"log"
	"strings"
)

// Define a callback function type that returns an error
type callback func([]string) ([]string, error)

// Step-by-step functions that compose the callback
func capitalizeWords(words []string) ([]string, error) {
	var result []string
	for _, word := range words {
		if len(word) == 0 {
			log.Printf("Error: Invalid input word '%s'", word)
			return nil, fmt.Errorf("invalid input: empty word '%s'", word)
		}
		result = append(result, strings.Title(word))
	}
	return result, nil
}

func removeDuplicates(words []string) ([]string, error) {
	var result []string
	seen := make(map[string]struct{})
	for _, word := range words {
		if _, ok := seen[word]; !ok {
			result = append(result, word)
			seen[word] = struct{}{}
		}
	}
	return result, nil
}

func composeCallbacks(callbacks ...callback) callback {
	return func(data []string) ([]string, error) {
		result := data
		for _, cb := range callbacks {
			var err error
			result, err = cb(result)
			if err != nil {
				log.Printf("Error in processing: %v", err)
				return nil, err
			}
		}
		return result, nil
	}
}

func main() {
	data := []string{"hello", "", "go", "language", "hello", "module", "", "golang"}

	callbackFunc := composeCallbacks(capitalizeWords, removeDuplicates)

	transformedData, err := callbackFunc(data)
	if err != nil {
		log.Fatalf("Failed to process data: %v", err)
	}

	fmt.Println(transformedData) // Output: [Hello Go Language Module Golang]
}