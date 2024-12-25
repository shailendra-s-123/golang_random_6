package main

import (
	"fmt"
	"strings"
	"time"
)

// Define a simple struct to represent a question
type Question struct {
	Text     string
	CorrectAnswer string
	Responses   []string
	ResponseTimes []time.Duration
	Skipped     []bool
}

// Calculate the total number of correct answers for all questions
func countCorrectAnswers(questions []Question) int {
	correctCount := 0
	for _, question := range questions {
		correctCount += countCorrectAnswersForQuestion(question)
	}
	return correctCount
}

// Calculate the number of correct answers for a single question
func countCorrectAnswersForQuestion(question Question) int {
	correctCount := 0
	for i, response := range question.Responses {
		if !question.Skipped[i] && strings.ToLower(response) == strings.ToLower(question.CorrectAnswer) {
			correctCount++
		}
	}
	return correctCount
}

// Identify the most common incorrect answers for all questions
func findMostCommonIncorrectAnswers(questions []Question) map[string]int {
	incorrectAnswers := make(map[string]int)

	for _, question := range questions {
		for i, response := range question.Responses {
			if !question.Skipped[i] && strings.ToLower(response) != strings.ToLower(question.CorrectAnswer) {
				incorrectAnswers[response]++
			}
		}
	}

	return incorrectAnswers
}

// Calculate the average response time for all questions
func calculateAverageResponseTime(questions []Question) time.Duration {
	totalTime := time.Duration(0)
	totalCount := 0

	for _, question := range questions {
		for _, responseTime := range question.ResponseTimes {
			if responseTime > 0 {
				totalTime += responseTime
				totalCount++
			}
		}
	}

	if totalCount == 0 {
		return 0
	}

	return totalTime / time.Duration(totalCount)
}

// Identify questions that were most frequently skipped
func findMostSkippedQuestions(questions []Question) map[int]int {
	skippedCounts := make(map[int]int)

	for i, question := range questions {
		skippedCounts[i] = 0
		for _, skipped := range question.Skipped {
			if skipped {
				skippedCounts[i]++
			}
		}
	}

	return skippedCounts
}

// Print insights from the data
func printInsights(questions []Question) {
	totalQuestions := len(questions)
	correctAnswers := countCorrectAnswers(questions)
	fmt.Printf("Total questions: %d\n", totalQuestions)
	fmt.Printf("Total correct answers: %d\n", correctAnswers)
	fmt.Printf("Average correct answers per question: %.2f\n", float64(correctAnswers)/float64(totalQuestions))

	incorrectAnswers := findMostCommonIncorrectAnswers(questions)
	fmt.Printf("Most common incorrect answers:\n")
	for answer, count := range incorrectAnswers {
		fmt.Printf("%s: %d times\n", answer, count)
	}

	averageResponseTime := calculateAverageResponseTime(questions)
	fmt.Printf("Average response time: %s\n", averageResponseTime)

	skippedCounts := findMostSkippedQuestions(questions)
	fmt.Printf("Most skipped questions:\n")
	for i, count := range skippedCounts {
		fmt.Printf("Question %d: %d times\n", i+1, count)
	}
}

func main() {
	// Sample questions
	questions := []Question{
		Question{
			Text:     "What is the capital of France?",
			CorrectAnswer: "Paris",
			Responses:   []string{"Paris", "Berlin", "Madrid", "Paris"},
			ResponseTimes: []time.Duration{3 * time.Second, 5 * time.Second, 4 * time.Second, 2 * time.Second},
			Skipped:     []bool{false, false, false, false},
		},
		Question{
			Text:     "What is 2 + 2?",
			CorrectAnswer: "4",
			Responses:   []string{"4", "2", "4", "5", "4"},
			ResponseTimes: []time.Duration{1 * time.Second, 3 * time.Second, 2 * time.Second, 4 * time.Second, 0 * time.Second},
			Skipped:     []bool{false, false, false, false, true},
		},
	}

	// Print insights
	printInsights(questions)
}