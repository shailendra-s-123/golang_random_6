package main

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

// StockPrice represents a stock price entry.
type StockPrice struct {
	Date  time.Time
	Price float64
}

// CalculateStats calculates the sum, average, and standard deviation of stock prices for a given period.
func CalculateStats(ctx context.Context, prices []StockPrice) (sum, average, stdDev float64, err error) {
	wg := sync.WaitGroup{}
	var results []float64

	wg.Add(3)
	// Calculate sum and average concurrently
	go func() {
		defer wg.Done()
		sum, average, err = calculateSumAndAverage(prices)
	}()

	// Calculate standard deviation concurrently
	go func() {
		defer wg.Done()
		results = calculateStandardDeviation(prices, average)
	}()

	wg.Wait()

	if err != nil {
		return 0, 0, 0, err
	}

	// Calculate standard deviation
	if len(results) > 0 {
		sumSquares := 0.0
		for _, result := range results {
			sumSquares += result * result
		}
		stdDev = math.Sqrt(sumSquares / float64(len(results)))
	}

	return
}

// calculateSumAndAverage calculates the sum and average of stock prices.
func calculateSumAndAverage(prices []StockPrice) (sum, average float64, err error) {
	if len(prices) == 0 {
		return 0, 0, errors.New("no prices to calculate")
	}

	var totalPrice float64
	for _, price := range prices {
		totalPrice += price.Price
	}

	return totalPrice, totalPrice / float64(len(prices)), nil
}

// calculateStandardDeviation calculates the standard deviation of stock prices.
func calculateStandardDeviation(prices []StockPrice, average float64) []float64 {
	var deviations []float64
	for _, price := range prices {
		deviation := price.Price - average
		deviations = append(deviations, deviation)
	}

	return deviations
}

// FetchData simulates fetching stock price data with a random delay.
func FetchData(ctx context.Context, numDays int) ([]StockPrice, error) {
	if numDays <= 0 {
		return nil, errors.New("numDays must be a positive integer")
	}

	// Simulate delay for data fetching
	delay := time.Duration(rand.Intn(5)) * time.Second
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(delay):
	}

	// Generate random stock prices for the specified number of days
	var prices []StockPrice
	currentDate := time.Now()
	for i := 0; i < numDays; i++ {
		price := rand.Float64() * 100 // Generate random price between 0 and 100
		prices = append(prices, StockPrice{Date: currentDate, Price: price})
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return prices, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	numDays := 10
	fmt.Printf("Fetching data for %d days...\n", numDays)

	// Fetch stock price data
	prices, err := FetchData(ctx, numDays)
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		return
	}

	// Calculate statistics
	sum, average, stdDev, err := CalculateStats(ctx, prices)