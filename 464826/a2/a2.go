package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// StockPrice represents a stock price entry.
type StockPrice struct {
	Date  time.Time
	Price float64
}

// FilterPrices filters stock prices based on a given minimum price.
func FilterPrices(ctx context.Context, prices []StockPrice, minPrice float64) []StockPrice {
	var filtered []StockPrice
	for _, price := range prices {
		select {
		case <-ctx.Done():
			return nil
		default:
			if price.Price >= minPrice {
				filtered = append(filtered, price)
			}
		}
	}
	return filtered
}

// AggregatePrices computes sum and average of filtered stock prices.
func AggregatePrices(ctx context.Context, prices []StockPrice) (sum, average float64) {
	totalPrice := 0.0
	count := 0

	for _, price := range prices {
		select {
		case <-ctx.Done():
			return 0, 0
		default:
			totalPrice += price.Price
			count++
		}
	}

	if count == 0 {
		return 0, 0
	}

	return totalPrice, totalPrice / float64(count)
}

// TransformPrices applies a transformation function to each price.
func TransformPrices(ctx context.Context, prices []StockPrice, transform func(float64) float64) []float64 {
	var transformed []float64

	var wg sync.WaitGroup
	chanSize := 100
	chanJobs := make(chan *StockPrice, chanSize)
	chanTransformed := make(chan float64, chanSize)

	defer close(chanJobs)
	defer close(chanTransformed)

	// Worker goroutine
	for {
		select {
		case item := <-chanJobs:
			defer wg.Done()
			price := item.Price
			select {
			case <-ctx.Done():
				return nil
			case chanTransformed <- transform(price):
			}
		default:
			if ctx.Err() != nil {
				return nil
			}
			return transformed
		}
	}

	// Main loop to send work
	for _, price := range prices {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case chanJobs <- &price:
			case <-ctx.Done():
				return
			}
		}()
	}

	wg.Wait()
	close(chanTransformed)

	for price := range chanTransformed {
		transformed = append(transformed, price)
	}

	return transformed
}

// FetchData simulates fetching stock price data.
func FetchData(ctx context.Context) ([]StockPrice, error) {
	// Simulate fetching data from a database or API
	dates := make([]time.Time, 1000000)
	prices := make([]float64, 1000000)
	for i := range dates {
		dates[i] = time.Date(2023, 10, 1, i/24, i%24, 0, 0, time.UTC)
		prices[i] = 100.0 + float64(i)/1000
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(100 * time.Millisecond):
	}

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

	// Filter prices over $120
	filteredPrices := FilterPrices(ctx, prices, 120.0)
	if filteredPrices == nil {
		fmt.Printf("Filtering canceled.\n")
		return
	}

	// Apply transformation: double each price
	transformedPrices := TransformPrices(ctx, filteredPrices, func(p float64) float64 {
		return p * 2
	})
	if transformedPrices == nil {
		fmt.Printf("Transformation canceled.\n")
		return
	}

	// Aggregate transformed prices
	sum, average := AggregatePrices(ctx, filteredPrices)

	// Display results
	fmt.Printf("Filtered prices count: %d\n", len(filteredPrices))
	fmt.Printf("First 10 filtered prices: %v\n", filteredPrices[:10])
	fmt.Printf("First 10 transformed prices: %v\n", transformedPrices[:10])
	fmt.Printf("Sum of filtered prices: %.2f\n", sum)
	fmt.Printf("Average of filtered prices: %.2f\n", average)
}