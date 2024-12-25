package main

import (
	"fmt"
	"strings"
	"time"
)

type Question struct {
	Text         string
	CorrectAnswer string
	Responses    []string
	ResponseTimes []time.Duration
	Skipped      []bool
}

// Function to calculate the average response time for a question
func averageResponseTime(question Question) time.Duration {
	totalTime := time.Duration(0)
	count := 0
	for i, response := range question.Responses {
		if !question.Skipped[i] {
			totalTime += question.ResponseTimes[i]
			count++
		}
	}
	if count == 0 {
		return time.Duration(0)
	}
	return totalTime / time.Duration(count)
}

// Function to find the question with the most skips
func mostSkippedQuestion(questions []Question) (string, int) {
	maxSkips := 0
	mostSkipped := ""
	for _, question := range questions {
		skips := 0
		for _, skipped := range question.Skipped {
			if skipped {
				skips++
			}
		}
		if skips > maxSkips {
			maxSkips = skips
			mostSkipped = question.Text
		}
	}
	return mostSkipped, maxSkips
}

func printInsights(questions []Question) {
	// Existing insights...
	fmt.Println("\n-- Additional Insights --")

	for i, question := range questions {
		fmt.Printf("\nQuestion %d: %s\n", i+1, question.Text)
		fmt.Printf("Average response time: %s\n", averageResponseTime(question))
	}

	mostSkipped, maxSkips := mostSkippedQuestion(questions)
	fmt.Printf("\nQuestion with the most skips: %s (Skipped %d times)\n", mostSkipped, maxSkips)
}

func main() {
	questions := []Question{
		// Existing questions...
	}

	// Collect responses and analyze additional data
	for i, question := range questions {
		fmt.Printf("Question %d: %s\n", i+1, question.Text)
		var response string
		fmt.Scanln(&response)

		startTime := time.Now()
		// Process user response and add it to the corresponding slices
		question.Responses = append(question.Responses, response)

		if strings.ToLower(response) == "skip" {
			question.Skipped = append(question.Skipped, true)
			question.ResponseTimes = append(question.ResponseTimes, time.Duration(0))
		} else {
			question.Skipped = append(question.Skipped, false)
			question.ResponseTimes = append(question.ResponseTimes, time.Since(startTime))
		}
	}

	// Print insights
	printInsights(questions)
}
 