package main  
import (  
    "fmt"
    "strings"
)
//Example task: Capitalize each word in the list and remove duplicates
func transformList(input []string) []string {  
    var result []string
    for _, word := range input {  
        capitalized := strings.Title(word)
        if !contains(result, capitalized) {
            result = append(result, capitalized)
        }
    }
    return result
}
func contains(slice []string, element string) bool {  
    for _, item := range slice {
        if item == element {
            return true
        }
    }
    return false
}
func main() {  
    data := []string{"hello", "go", "language", "hello", "module"}
    transformedData := transformList(data)  
    fmt.Println(transformedData) // Output: [Hello Go Language Module]
}