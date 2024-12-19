package main  
import (  
    "fmt"
    "strings"
)

// Define a UserResponse struct to store user response for each question
type UserResponse struct {
    Question   string
    Answer     string
    IsCorrect bool
}

func main() {  
    // Sample questions
    questions := []string{
        "What is the capital of France?",
        "What is 2 + 2?",
        "What is the largest ocean?",
    }
    // Sample correct answers
    correctAnswers := []string{"Paris", "4", "Pacific Ocean"}
    var responses []UserResponse
    
    // Simulate user responses
    userResponses := [][]string{
        {"Paris", "2", "Atlantic Ocean"},
        {"Berlin", "4", "Pacific Ocean"},
        {"Paris", "3", "Pacific Ocean"},
        {"Berlin", "5", "Indian Ocean"},
        {"Paris", "4", "Atlantic Ocean"},
    }
    
    // Process user responses and store in slices
    for i, userAnswer := range userResponses {
        for j, ans := range userAnswer {
            isCorrect := strings.ToLower(ans) == strings.ToLower(correctAnswers[j])
            responses = append(responses, UserResponse{
                Question: questions[j],
                Answer:   ans,
                IsCorrect: isCorrect,
            })
        }
    }
    
    // Aggregate data and calculate insights
    correctCount, responseFrequencies, mostCommonIncorrect := aggregateData(responses)
    userPerformance := calculateUserPerformance(correctCount, len(responses))

    // Display insights
    fmt.Println("Analysis Insights:")
    fmt.Println("--------------------")
    fmt.Printf("Total correct answers: %d\n", correctCount)
    fmt.Println("Response Frequencies:")
    for answer, freq := range responseFrequencies {
        fmt.Printf("%s: %d\n", answer, freq)
    }
    fmt.Println("Most common incorrect answers:")
    for _, incorrect := range mostCommonIncorrect {
        fmt.Println(incorrect)
    }
    fmt.Printf("User Performance: %.2f%%\n", userPerformance)
}
 
func aggregateData(responses []UserResponse) (int, map[string]int, []string) {
    correctCount := 0
    responseFrequencies := make(map[string]int)
    incorrectAnswers := make(map[string]int)
    mostCommonIncorrect := []string{}

    for _, response := range responses {
        if response.IsCorrect {
            correctCount++
        }
        responseFrequencies[response.Answer]++
        if !response.IsCorrect {
            incorrectAnswers[response.Answer]++
        }
    }
    
    // Find the most common incorrect answer(s)
    maxIncorrectCount := 0
    for answer, count := range incorrectAnswers {
        if count > maxIncorrectCount {
            mostCommonIncorrect = []string{answer}
            maxIncorrectCount = count
        } else if count == maxIncorrectCount {
            mostCommonIncorrect = append(mostCommonIncorrect, answer)
        }
    }

    return correctCount, responseFrequencies, mostCommonIncorrect
}

func calculateUserPerformance(correctCount int, totalResponses int) float64 {
    return float64(correctCount) / float64(totalResponses) * 100
} 