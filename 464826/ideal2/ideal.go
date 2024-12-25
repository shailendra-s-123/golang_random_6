package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Transaction represents a financial transaction
type Transaction struct {
	ID     int
	Amount float64
	Type   string // e.g., "credit" or "debit"
}

// FilterCriteria defines the criteria for filtering transactions
type FilterCriteria struct {
	MinAmount float64
	Type      string
}

func main() {
	// Generate a large dataset of transactions
	transactions := generateTransactions(1_000_000)

	// Define filter criteria
	criteria := FilterCriteria{
		MinAmount: 100.0,
		Type:      "credit",
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Process the transactions concurrently
	filteredTransactions, err := processTransactions(ctx, transactions, criteria)
	if err != nil {
		fmt.Printf("Error processing transactions: %v\n", err)
		return
	}

	// Display results
	fmt.Printf("Processed %d transactions meeting criteria.\n", len(filteredTransactions))
	fmt.Printf("Total amount: %.2f\n", aggregateAmount(filteredTransactions))
}

// generateTransactions generates a dataset of random transactions
func generateTransactions(n int) []Transaction {
	rand.Seed(time.Now().UnixNano())
	types := []string{"credit", "debit"}
	transactions := make([]Transaction, n)
	for i := range transactions {
		transactions[i] = Transaction{
			ID:     i + 1,
			Amount: rand.Float64() * 1000,
			Type:   types[rand.Intn(len(types))],
		}
	}
	return transactions
}

// processTransactions filters and transforms transactions concurrently
func processTransactions(ctx context.Context, transactions []Transaction, criteria FilterCriteria) ([]Transaction, error) {
	var wg sync.WaitGroup
	results := make(chan Transaction, len(transactions))
	errCh := make(chan error, 1)

	// Divide transactions into chunks for concurrent processing
	chunkSize := 10_000
	for start := 0; start < len(transactions); start += chunkSize {
		end := start + chunkSize
		if end > len(transactions) {
			end = len(transactions)
		}

		wg.Add(1)
		go func(chunk []Transaction) {
			defer wg.Done()
			for _, t := range chunk {
				select {
				case <-ctx.Done():
					errCh <- ctx.Err()
					return
				default:
					if t.Amount >= criteria.MinAmount && t.Type == criteria.Type {
						results <- t
					}
				}
			}
		}(transactions[start:end])
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(results)
		close(errCh)
	}()

	// Collect results or handle errors
	var filtered []Transaction
	for {
		select {
		case t, ok := <-results:
			if !ok {
				return filtered, nil
			}
			filtered = append(filtered, t)
		case err := <-errCh:
			return nil, err
		}
	}
}

// aggregateAmount calculates the total amount of filtered transactions
func aggregateAmount(transactions []Transaction) float64 {
	var total float64
	for _, t := range transactions {
		total += t.Amount
	}
	return total
}
