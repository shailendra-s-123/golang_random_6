package main  
import (  
    "context"  
    "fmt"  
    "sort"  
    "sync"  
)  
// FinancialData represents a single record of financial data.  
type FinancialData struct {  
    Date string  
    Open float64  
    High float64  
    Low float64  
    Close float64  
    Volume float64  
}  
// ProcessFinancialData takes a slice of FinancialData as input, calculates some   
// technical indicators, and prints the results to the console.  
func ProcessFinancialData(ctx context.Context, data []FinancialData) error {  
    // Return an error if the context is cancelled.  
    select {  
    case <-ctx.Done():  
        return ctx.Err()  
    default:  
    }  
    // Calculate the Simple Moving Average (SMA) over a window of 5 days.  
    sma := calculateSMA(data, 5)  
    // Calculate the Exponential Moving Average (EMA) over a window of 10 days.  
    ema := calculateEMA(data, 10)  
    // Calculate the Relative Strength Index (RSI) over a window of 14 days.  
    rsi := calculateRSI(data, 14)  
    // Print the results.  
    fmt.Println("Simple Moving Average (SMA):")  
    for i, value := range sma {  
        fmt.Printf("Date: %s, SMA: %f\n", data[i].Date, value)  
    }  
    fmt.Println("\nExponential Moving Average (EMA):")  
    for i, value := range ema {  
        fmt.Printf("Date: %s, EMA: %f\n", data[i].Date, value)  
    }  
    fmt.Println("\nRelative Strength Index (RSI):")  
    for i, value := range rsi {  
        fmt.Printf("Date: %s, RSI: %f\n", data[i].Date, value)  
    }  
    return nil  
}  
func calculateSMA(data []FinancialData, window int) []float64 {  
    // Return an empty slice if the window size is greater than the data size.  
    if window > len(data) {  
        return []float64{}  
    }  
    sma := make([]float64, len(data))  
    for i := window - 1; i < len(data); i++ {  
        sum := float64(0)  
        for j := i - window + 1; j <= i; j++ {  
            sum += data[j].Close  
        }  
        sma[i] = sum / float64(window)  
    }  
    return sma  
}  
func calculateEMA(data []FinancialData, window int) []float64 {  
    // Return an empty slice if the window size is greater than the data size.  
    if window > len(data) {  
        return []float64{}  
    }  
    ema := make([]float64, len(data))  
    // Initialize the EMA for the first window elements.  
    for i := 0; i < window; i++ {  
        ema[i] = data[i].Close  
    }  
    // Calculate the EMA for the remaining elements.  
    smoothingFactor := 2.0 / float64(window + 1)  
    for i := window; i < len(data); i++ {  
        ema[i] = (data[i].Close - ema[i-1]) * smoothingFactor + ema[i-1]  
    }  
    return ema  
}  
func calculateRSI(data []FinancialData, window int) []float64 {  
    // Return an empty slice if the window size is greater than the data size.  
    if window > len(data) {  
        return []float64{}  
    }  
    rsi := make([]float64, len(data))  
    gains := make([]float64, window)  