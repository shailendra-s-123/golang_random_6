package main

import (
	"fmt"
	"sort"
	"strings"
)

// Define a simple struct to represent a question
type Question struct {
	Text     string
	CorrectAnswer string
	Responses   []string
}

// Function to calculate response frequencies
func calculateResponseFrequencies(responses []string) map[string]int {
	freqs := make(map[string]int)
	for _, response := range responses {
		freqs[response]++
	}
	return freqs
}

// Function to find the most common incorrect answers
func findMostCommonIncorrectAnswers(questions []Question) map[string]int {
	incorrectAnswers := make(map[string]int)

	for _, question := range questions {
		correctAnswer := strings.ToLower(question.CorrectAnswer)
		for _, response := range question.Responses {
			if strings.ToLower(response) != correctAnswer {
				incorrectAnswers[response]++
			}
		}
	}

	// Sort by count in descending order
	type answerWithCount struct {
		answer    string
		count     int
	}
	var answerCounts []answerWithCount
	for answer, count := range incorrectAnswers {
		answerCounts = append(answerCounts, answerWithCount{answer, count})
	}
	sort.Slice(answerCounts, func(i, j int) bool {
		return answerCounts[i].count > answerCounts[j].count
	})

	mostCommon := make(map[string]int)
	for _, ac := range answerCounts[:3] { // Get top 3 most common incorrect answers
		mostCommon[ac.answer] = ac.count
	}
	return mostCommon
}

// Function to calculate user performance in percentage
func calculateUserPerformance(questions []Question) map[string]float64 {
	performance := make(map[string]float64)
	totalUsers := len(questions[0].Responses)
	totalQuestions := len(questions)

	// Assuming all questions have the same number of users' responses
	for _, question := range questions {
		correctCount := countCorrectAnswersForQuestion(question)
		for i, response := range question.Responses {
			if strings.ToLower(response) == strings.ToLower(question.CorrectAnswer) {
				userID := fmt.Sprintf("user%d", i+1)
				performance[userID] = calculatePerformanceForUser(performance, userID, correctCount)
			}
		}
	}

	return performance
}

// Function to count correct answers for a single question
func countCorrectAnswersForQuestion(question Question) int {
	correctCount := 0
	for _, response := range question.Responses {
		if strings.ToLower(response) == strings.ToLower(question.CorrectAnswer) {
			correctCount++
		}
	}
	return correctCount
}

// Function to calculate performance for a user
func calculatePerformanceForUser(performance map[string]float64, userID string, correctAnswers int) float64 {
	if _, ok := performance[userID]; !ok {
		performance[userID] = 0.0
	}
	performance[userID] += float64(correctAnswers) / float64(totalQuestions) * 100.0
	return performance[userID]
}

// Print insights from the data
func printInsights(questions []Question) {
	fmt.Printf("\nResponse Frequencies:\n")
	for _, question := range questions {
		freqs := calculateResponseFrequencies(question.Responses)
		for answer, count := range freqs {
			fmt.Printf("  %s: %d\n", answer, count)
		}
	}

	fmt.Printf("\nMost Common Incorrect Answers:\n")
	incorrectAnswers := findMostCommonIncorrectAnswers(questions)
	for answer, count := range incorrectAnswers {
		fmt.Printf("  %s: %d\n", answer, count)
	}

	fmt.Printf("\nUser Performance (\\%):\n")
	performance := calculateUserPerformance(questions)
	for userID, score := range performance {
		fmt.Printf("  %s: %.2f\n", userID, score)
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