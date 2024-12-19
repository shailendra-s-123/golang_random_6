package main

import (
	"context"
	"fmt"
	_"math"
	"time"
)

// StockPrice represents a stock price entry.
type StockPrice struct {
	Date  time.Time
	Price float64
}

// CalculateStats calculates the sum and average of stock prices for a given period.
func CalculateStats(ctx context.Context, prices []StockPrice) (sum, average float64) {
	totalPrice := 0.0
	count := 0

	for _, price := range prices {
		select {
		case <-ctx.Done():
			return 0, 0 // Cancel operation gracefully
		default:
			totalPrice += price.Price
			count++
		}
	}

	if count == 0 {
		return 0, 0 // Avoid division by zero
	}

	return totalPrice, totalPrice / float64(count)
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
	}

	prices := []float64{100.0, 101.0, 102.0}

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
	sum, average := CalculateStats(ctx, prices)

	// Display results
	fmt.Printf("Sum of prices: %.2f\n", sum)
	fmt.Printf("Average price: %.2f\n", average)
}