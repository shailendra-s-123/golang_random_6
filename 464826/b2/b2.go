package main  
import (  
    "context"
    "fmt"
    "math"
    "runtime"
    "sync"
    "time"
)

// StockPrice represents a stock price entry.
type StockPrice struct {
    Date  time.Time
    Price float64
}

// CalculateStats calculates the sum, average, and median of stock prices for a given period.
func CalculateStats(ctx context.Context, prices []StockPrice, wg *sync.WaitGroup, results chan<- map[string]float64) {
    defer wg.Done()

    totalPrice := 0.0
    count := 0
    priceMap := make(map[float64]int)

    for _, price := range prices {
        select {
        case <-ctx.Done():
            return // Cancel operation gracefully
        default:
            totalPrice += price.Price
            count++
            priceMap[price.Price]++
        }
    }

    if count == 0 {
        results <- nil
        return
    }

    // Calculate median
    pricesSorted := make([]float64, 0, count)
    for price, frequency := range priceMap {
        for i := 0; i < frequency; i++ {
            pricesSorted = append(pricesSorted, price)
        }
    }
    median := calculateMedian(pricesSorted)

    // Send results through the channel
    results <- map[string]float64{
        "Sum":    totalPrice,
        "Average": totalPrice / float64(count),
        "Median":  median,
    }
}

func calculateMedian(prices []float64) float64 {
    length := len(prices)
    if length == 0 {
        return 0
    }

    if length%2 == 0 {
        mid1 := length / 2 - 1
        mid2 := length / 2
        return (prices[mid1] + prices[mid2]) / 2
    } else {
        mid := length / 2
        return prices[mid]
    }
}

// FilterPrices filters stock prices based on a given threshold.
func FilterPrices(ctx context.Context, prices []StockPrice, threshold float64) []StockPrice {
    var filteredPrices []StockPrice

    for _, price := range prices {
        select {
        case <-ctx.Done():
            return nil // Cancel operation gracefully
        default:
            if price.Price > threshold {
                filteredPrices = append(filteredPrices, price)
            }
        }
    }

    return filteredPrices
}

// AggregatePrices aggregates stock prices into daily buckets.
func AggregatePrices(ctx context.Context, prices []StockPrice) map[time.Time]float64 {
    priceBuckets := make(map[time.Time]float64)

    for _, price := range prices {
        select {
        case <-ctx.Done():
            return nil // Cancel operation gracefully
        default:
            priceBuckets[price.Date.Truncate(24*time.Hour)] += price.Price
        }
    }

    return priceBuckets
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