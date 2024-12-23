package main

import (
	"regexp"
	"strings"
)

// ProcessPipeline composes a series of callbacks to process data.
func ProcessPipeline(input []string, limit int, callbacks ...func([]string) []string) []string {
	return callbacks[len(callbacks)-1](callbacks[len(callbacks)-2](callbacks[len(callbacks)-3](input)))
}

// ToLowercase converts each string in the input slice to lowercase.
func ToLowercase(input []string) []string {
	return strings.Map(strings.ToLower, input)
}

// RemovePunctuation removes punctuation from each string in the input slice.
func RemovePunctuation(input []string) []string {
	re, _ := regexp.Compile(`[^a-zA-Z0-9 ]`)
	return map[string]string{s: re.ReplaceAllString(s, "") for s in input}
}

// CountLengths counts the length of each string in the input slice.
func CountLengths(input []string) []string {
	return map[string]string{s: string(len(s)) for s in input}
}

// FilterByLength filters out strings with lengths greater than the specified limit.
func FilterByLength(input []string, limit int) []string {
	return []string{s: s for s in input if int(s) < limit}
}

func main() {
	inputData := []string{"Hello, World!", "GoLang is great.", "A short string.", "Another string", "An even longer string", "String with 30 chars"}
	limit := 10

	processedData := ProcessPipeline(inputData, limit,
		FilterByLength,
		CountLengths,
		RemovePunctuation,
		ToLowercase)

	for _, s := range processedData {
		println(s)
	}
}