package main

import (
"context"
"fmt"
"time"
)

// StockPrice represents a stock price entry with a date and a price.
type StockPrice struct {
Date time.Time
Price float64
}

// FetchData simulates fetching stock price data. It uses the context to handle cancellations.
func FetchData(ctx context.Context) ([]StockPrice, error) {
// Simulate fetching data from a database or external API. Adding a delay for demonstration.
select {
case <-ctx.Done():
// If the context is done, return an error.
return nil, ctx.Err()
case <-time.After(2 * time.Second):
// Simulating a successful fetch with 2 seconds delay.
}

// Example data with 3 dates and corresponding prices.
dates := []time.Time{
time.Date(2023, 7, 28, 0, 0, 0, 0, time.UTC),
time.Date(2023, 7, 29, 0, 0, 0, 0, time.UTC),
time.Date(2023, 7, 30, 0, 0, 0, 0, time.UTC),
}

prices := []float64{100.0, 101.5, 102.0}

// Combine the date and price data into a slice of StockPrice.
return zipData(dates, prices), nil
}

// zipData combines two slices into a slice of StockPrice structs.
func zipData(dates []time.Time, prices []float64) []StockPrice {
var data []StockPrice
// Iterate over dates slice and combine with prices.
for i := range dates {
if i < len(prices) {
// Create StockPrice instances and append to the data slice.
data = append(data, StockPrice{dates[i], prices[i]})
}
}
return data
}

// CalculateStats calculates the sum and average of stock prices. It also respects the context for handling cancellations.
func CalculateStats(ctx context.Context, prices []StockPrice) (sum, average float64) {
totalPrice := 0.0
count := 0

// Iterating over the stock prices while checking for cancellation.
for _, price := range prices {
select {
case <-ctx.Done():
// If the context is canceled, gracefully return.
return 0, 0
default:
// Accumulate total price and count of stock prices.
totalPrice += price.Price
count++
}
}

// Avoid division by zero if no prices are present.
if count == 0 {
return 0, 0
}

// Return the sum and average of the stock prices.
return totalPrice, totalPrice / float64(count)
}

// calculateTrends calculates the trend of stock prices by comparing consecutive prices.
func calculateTrends(prices []StockPrice) []string {
trends := make([]string, 0, len(prices)-1)
for i := 1; i < len(prices); i++ {
if prices[i].Price > prices[i-1].Price {
trends = append(trends, "Up")
} else if prices[i].Price < prices[i-1].Price {
trends = append(trends, "Down")
} else {
trends = append(trends, "Flat")
}
}
return trends
}

func main() {
// Create a context with a timeout of 5 seconds.
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel() // Ensure cancellation of the context once the main function completes.

// Fetch the stock price data.
prices, err := FetchData(ctx)
if err != nil {
// Handle error if data fetching fails or is canceled.
fmt.Printf("Error fetching data: %v\n", err)
return