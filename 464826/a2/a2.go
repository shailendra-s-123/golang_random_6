package main

import (
	"context"
	"fmt"
	"math"
	"time"
)

// StockPrice represents a stock price entry.
type StockPrice struct {
	Date  time.Time
	Price float64
}

// CalculateStats calculates the sum, average, and trend of stock prices.
func CalculateStats(ctx context.Context, prices []StockPrice) (sum, average, trend float64) {
	totalPrice := 0.0
	count := 0
	prevPrice := 0.0

	for _, price := range prices {
		select {
		case <-ctx.Done():
			return 0, 0, 0 // Cancel operation gracefully
		default:
			totalPrice += price.Price
			count++

			if count > 1 {
				trend = price.Price - prevPrice
			}

			prevPrice = price.Price
		}
	}

	if count == 0 {
		return 0, 0, 0 // Avoid division by zero
	}

	return totalPrice, totalPrice / float64(count), trend
}

// FetchData simulates fetching stock price data.
func FetchData(ctx context.Context) ([]StockPrice, error) {
	// Simulate fetching data from a database or API
	// Delay for demonstration purposes
	select {
	case <-ctx.Done():
		return nil, ctx.Err() // Cancel data fetching
	case <-time.After(2 * time.Second):
	}

	// Example data
	dates := []time.Time{
		time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 10, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 10, 4, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 10, 5, 0, 0, 0, 0, time.UTC),
	}

	prices := []float64{100.0, 101.0, 102.0, 103.0, 104.0}

	return zipData(dates, prices), nil
}

// zipData combines two slices into a slice of StockPrice.
func zipData(dates []time.Time, prices []float64) []StockPrice {
	var data []StockPrice
	for i := range dates {
		if i < len(prices) {
			data = append(data, StockPrice{dates[i], prices[i]})
		}
	}
	return data
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Simulate fetching stock price data
	prices, err := FetchData(ctx)
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		return
	}

	// Calculate statistics
	sum, average, trend := CalculateStats(ctx, prices)

	// Display results
	fmt.Printf("Sum of prices: %.2f\n", sum)
	fmt.Printf("Average price: %.2f\n", average)
	fmt.Printf("Trend of prices: %.2f\n", trend)

	// Suggestion: Extend the program to calculate moving averages
	fmt.Println("\nSuggestion: Extend the program to calculate moving averages.")
	fmt.Println("You can implement a function to calculate N-day moving averages.")
	fmt.Println("Example: Calculate 3-day moving average for the given prices.")
}