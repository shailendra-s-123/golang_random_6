package main
import ("fmt"
)
type UserResponse struct {
    UserID    string
    QuestionID int
    Answer    string
    IsCorrect bool
}
func main() {
    // Sample user responses
    responses := []UserResponse{
        {UserID: "user1", QuestionID: 1, Answer: "A", IsCorrect: true},
        {UserID: "user1", QuestionID: 2, Answer: "B", IsCorrect: false},
        {UserID: "user2", QuestionID: 1, Answer: "A", IsCorrect: true},
        {UserID: "user2", QuestionID: 2, Answer: "C", IsCorrect: false},
        {UserID: "user3", QuestionID: 1, Answer: "B", IsCorrect: false},
        {UserID: "user3", QuestionID: 2, Answer: "B", IsCorrect: true},
    }
    // Calculate response frequencies for each answer option
    answerFrequencies := make(map[string]int)
    for _, response := range responses {
        answerFrequencies[response.Answer]++
    }
    // Identify popular incorrect answers
    incorrectAnswers := make(map[string]int)
    for _, response := range responses {
        if !response.IsCorrect {
            incorrectAnswers[response.Answer]++
        }
    }
    popularIncorrectAnswers := make([]string, 0)
    for answer, count := range incorrectAnswers {
        if count > 1 {
            popularIncorrectAnswers = append(popularIncorrectAnswers, answer)
        }
    }
    // Calculate user performance (percentage of correct answers)
    userPerformance := make(map[string]float64)
    for _, response := range responses {
        userPerformance[response.UserID] += float64(response.IsCorrect)
    }
    for userID := range userPerformance {
        userPerformance[userID] = (userPerformance[userID] / float64(len(responses))) * 100
    }
    // Output insights
    fmt.Println("Response Frequencies:")
    for answer, frequency := range answerFrequencies {
        fmt.Printf("Answer %s: %d times\n", answer, frequency)
    }
    fmt.Println("\nPopular Incorrect Answers:")
    for _, answer := range popularIncorrectAnswers {
        fmt.Printf("Answer %s: %d times\n", answer, incorrectAnswers[answer])
    }
    fmt.Println("\nUser Performance:")
    for userID, performance := range userPerformance {
        fmt.Printf("User %s: %.2f%% correct\n", userID, performance)
    }
}