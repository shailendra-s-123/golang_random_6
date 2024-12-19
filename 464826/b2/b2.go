package main

import (
	"context"
	"fmt"
	_"math"
	"time"
)

// StockPrice represents a stock price entry
type StockPrice struct {
	Date  time.Time
	Price float64
}

// calculateStats calculates the sum and average of stock prices.
func calculateStats(ctx context.Context, prices []StockPrice) (sum, average float64) {
	for _, price := range prices {
		select {
		case <-ctx.Done():
			return 0, 0 // Cancel operation gracefully
		default:
			sum += price.Price
		}
	}
	if len(prices) == 0 {
		return 0, 0
	}
	return sum, sum / float64(len(prices))
}

// fetchData simulates fetching stock price data.
func fetchData(ctx context.Context) ([]StockPrice, error) {
	// Simulate fetching data from a database or API.
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(2 * time.Second):
	}
	dates := []time.Time{
		time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 10, 3, 0, 0, 0, 0, time.UTC),
	}
	prices := []float64{100.0, 101.0, 102.0}
	return zipData(dates, prices), nil
}

// zipData combines dates and prices into StockPrice slice.
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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	prices, err := fetchData(ctx)
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		return
	}
	sum, average := calculateStats(ctx, prices)
	fmt.Printf("Sum of prices: %.2f\n", sum)
	fmt.Printf("Average price: %.2f\n", average)
}  