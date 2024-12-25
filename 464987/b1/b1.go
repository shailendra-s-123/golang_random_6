package main  
import (  
    "fmt"
    "log"
    "strings"
)

// Custom error type to distinguish between scenarios
type DashError struct {
    message string
    cause   error
}

func (e DashError) Error() string {
    return fmt.Sprintf("%s: %s", e.message, e.cause)
}

// Custom error for index out of bound
func IndexOutOfBoundError(message string, cause error) error {
    return DashError{message: message, cause: cause}
}

// Custom error for nil slice
func NilSliceError(message string) error {
    return DashError{message: message, cause: nil}
}
func fetchDataFromDatabase() ([]int, error) {  
    data := []int{1, 2, 3}
    // Simulate an error condition: accessing index out of bounds
    return data[:len(data)+1], IndexOutOfBoundError("fetchDataFromDatabase: Slice index out of bounds", nil)
}

func calculateSliceSum(data []int) (int, error) {  
    if data == nil {
        return 0, NilSliceError("calculateSliceSum: Received nil slice")
    }
    total := 0
    for _, value := range data {  
        total += value
    }
    return total, nil
}

func displayDashboard(data []int, total int) error {  
    // Handle nil slice gracefully
    if data == nil {
        log.Println("displayDashboard: Failed to calculate slice sum. Received nil slice.")
        fmt.Println("Dashboard: An error occurred. Please try again later.")
        return NilSliceError("displayDashboard: Failed to calculate slice sum. Received nil slice.")
    }
    // Handle other errors
    if _, err := calculateSliceSum(data); err != nil {
        wrappedErr := fmt.Errorf("displayDashboard: Failed to calculate slice sum: %w", err)
        log.Println(wrappedErr)
        fmt.Println("Dashboard: An error occurred. Please try again later.")  return wrappedErr
    }

    fmt.Println("Dashboard:")  
    fmt.Println("--------------")  
    fmt.Println("Slice Data:", data)  
    fmt.Println("Total:", total)
    fmt.Println("--------------")  
    return nil
}  

func main() {  
    data, err := fetchDataFromDatabase()  
    if err != nil {
        wrappedErr := fmt.Errorf("main: Failed to fetch data from database: %w", err)
        log.Println(wrappedErr)
        fmt.Println("An error occurred while fetching data. Please try again later.")
        // You can add more error handling based on the custom error types
        if strings.Contains(wrappedErr.Error(), "index out of bounds") {
            fmt.Println("Hint: Please check your slice index ranges.")
        }
        return
    }
    
    total, err := calculateSliceSum(data)  
    if err != nil {
        log.Println("main: Failed to calculate slice sum:", err)
        fmt.Println("An error occurred while calculating the total. Please try again later.")
        return
    }

    err = displayDashboard(data, total)  
    if err != nil {
        log.Println("main: Failed to display dashboard:", err)  
        fmt.Println("An error occurred while displaying the dashboard. Please try again later.")
    }  
}  