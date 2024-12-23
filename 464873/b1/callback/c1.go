package main  
import (  
    "fmt"
    "strings"
)
//Define a callback function type
type callback func([]string) []string

//Step-by-step functions that compose the callback
func capitalizeWords(words []string) []string {  
    var result []string
    for _, word := range words {
        result = append(result, strings.Title(word))  
    }
    return result
}
func removeDuplicates(words []string) []string {  
    var result []string
    seen := make(map[string]struct{})
    for _, word := range words {  
        if _, ok := seen[word]; !ok {  
            result = append(result, word)  
            seen[word] = struct{}{}  
        }  
    }  
    return result  
}  
func composeCallbacks(callbacks ...callback) callback {  
    return func(data []string) []string {  
        result := data
        for _, cb := range callbacks {  
            result = cb(result)  
        }  
        return result  
    }  
}
func main() {  
    data := []string{"hello", "go", "language", "hello", "module"}
    callbackFunc := composeCallbacks(capitalizeWords, removeDuplicates)
    transformedData := callbackFunc(data)
    fmt.Println(transformedData) // Output: [Hello Go Language Module]
}