package main

import (
	"fmt"
	"strings"
)

// Define a simple struct to represent a question
type Question struct {
	Text     string
	CorrectAnswer string
	Responses   []string
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
	for _, response := range question.Responses {
		if strings.ToLower(response) == strings.ToLower(question.CorrectAnswer) {
			correctCount++
		}
	}
	return correctCount
}

// Identify the most common incorrect answers for all questions
func findMostCommonIncorrectAnswers(questions []Question) map[string]int {
	incorrectAnswers := make(map[string]int)

	for _, question := range questions {
		for _, response := range question.Responses {
			if strings.ToLower(response) != strings.ToLower(question.CorrectAnswer) {
				incorrectAnswers[response]++
			}
		}
	}

	return incorrectAnswers
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
}

func main() {
	// Sample questions
	questions := []Question{
		Question{
			Text:     "What is the capital of France?",
			CorrectAnswer: "Paris",
			Responses:   []string{"Paris", "Berlin", "Madrid", "Paris", "Berlin"},
		},
		Question{
			Text:     "What is 2 + 2?",
			CorrectAnswer: "4",
			Responses:   []string{"4", "2", "4", "5", "4"},
		},
	}

	// Print insights
	printInsights(questions)
}