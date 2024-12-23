package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// StockPrice represents a stock price entry with date and price.
type StockPrice struct {
	Date  time.Time
	Price float64
}

// FetchData simulates fetching stock prices from an external source.
func FetchData(ctx context.Context) ([]StockPrice, error) {
	// Simulate fetching data from a database or API with a delay
	select {
	case <-ctx.Done():
		return nil, ctx.Err() // Return early if canceled
	case <-time.After(3 * time.Second):
	}

	// Example financial data
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
	for i, date := range dates {
		if i < len(prices) {
			data = append(data, StockPrice{date, prices[i]})
		}
	}
	return data
}

// CalculateStats calculates statistics on a set of stock prices.
func CalculateStats(ctx context.Context, prices []StockPrice) (sum, average float64, err error) {
	totalPrice := 0.0
	count := 0

	for _, price := range prices {
		select {
		case <-ctx.Done():
			return 0, 0, ctx.Err() // Handle cancelation
		default:
			totalPrice += price.Price
			count++
		}
	}

	if count == 0 {
		err = fmt.Errorf("no price data")
		return
	}

	return totalPrice, totalPrice / float64(count), nil
}

// DetectTrends detects trends (upward or downward) in stock prices.
func DetectTrends(ctx context.Context, prices []StockPrice) (upWard bool, err error) {
	if len(prices) < 2 {
		err = fmt.Errorf("not enough data to detect trends")
		return
	}

	previousPrice := prices[0].Price
	for _, price := range prices[1:] {
		select {
		case <-ctx.Done():
			return false, ctx.Err() // Handle cancelation
		default:
			if price.Price > previousPrice {
				upWard = true
			} else {
				upWard = false
				break
			}
			previousPrice = price.Price
		}
	}

	return upWard, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Fetch data with a chance to handle timeouts or cancelations
	stockPrices, err := FetchData(ctx)
	if err != nil {
		log.Printf("Error fetching data: %v", err)
		return
	}

	// Calculate statistics
	sum, average, err := CalculateStats(ctx, stockPrices)
	if err != nil {
		log.Printf("Error calculating stats: %v", err)
		return
	}

	// Detect trend in stock prices
	upWard, err := DetectTrends(ctx, stockPrices)
	if err != nil {
		log.Printf("Error detecting trend: %v", err)
		return
	}

	// Display results
	fmt.Printf("Sum of prices: %.2f\n", sum)
	fmt.Printf("Average price: %.2f\n", average)
	if upWard {
		fmt.Println("Trend: Upward")
	} else {
		fmt.Println("Trend: Downward")
	}
}