package main  
import (  
    "fmt"
    "strings"
)

// Define a callback function type with error return
type callback func([]string) ([]string, error)

//Step-by-step functions that compose the callback, now returning errors
func capitalizeWords(words []string) ([]string, error) {  
    var result []string  
    for _, word := range words {  
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
//** new function **
func divideByZero(numbers []int) ([]int, error) {
    var result []int
    for _, num := range numbers {
        if num == 0 {
            return result, fmt.Errorf("division by zero error")
        }
        result = append(result, num/0)
    }
    return result, nil
}

func composeCallbacks(callbacks ...callback) callback {  
    return func(data interface{}) ([]string, error) {  
        result := data
        for _, cb := range callbacks {  
            var err error
            result, err = cb(result.([]string)) 
            if err != nil {
                return nil, err
            }  
        }  
        return result.([]string), nil  
    }  
}
func main() {  
    data := []string{"hello", "go", "language", "hello", "module"}
    callbackFunc := composeCallbacks(capitalizeWords, removeDuplicates, divideByZero)
    transformedData, err := callbackFunc(data)  
    if err != nil {  
        fmt.Println("Error:", err)  
    } else {  
        fmt.Println(transformedData)
    }
}