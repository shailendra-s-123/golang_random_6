package main

import (
    "fmt"
    "sync"
)

const (
    // Buffer size for data channels
    channelBufferSize = 100
)

// Data type for the pipeline
type Data int

// ProcessFunc defines the signature for data processing functions
type ProcessFunc func(Data) Data

// CallbackChain represents the entire callback chain in the pipeline
type CallbackChain struct {
    processors []ProcessFunc
}

// NewCallbackChain returns a new CallbackChain
func NewCallbackChain(processors ...ProcessFunc) *CallbackChain {
    return &CallbackChain{processors: processors}
}

// Run executes the callback chain on the given data
func (c *CallbackChain) Run(data Data) Data {
    for _, processor := range c.processors {
        data = processor(data)
    }
    return data
}

// BackpressureConsumer handles data consumption with backpressure
func BackpressureConsumer(wg *sync.WaitGroup, input <-chan Data, output chan<- Data, chain *CallbackChain) {
    defer wg.Done()

    for data := range input {
        // Process the data with the callback chain
        processedData := chain.Run(data)

        // Send the processed data to the output channel
        output <- processedData
    }
}

// BackpressureProducer generates data with backpressure
func BackpressureProducer(wg *sync.WaitGroup, input chan<- Data, done <-chan bool) {
    defer wg.Done()

    defer close(input)

    for i := 1; ; i++ {
        // Generate a new data item
        data := Data(i)

        // Non-blocking send to the input channel with a capacity check
        select {
        case input <- data:
        case <-done:
            // Consumer has signaled done, so stop generating data
            return
        }
    }
}

func main() {
    // Initialize wait group for synchronization
    var wg sync.WaitGroup

    // Create data channels with buffering
    input := make(chan Data, channelBufferSize)
    output := make(chan Data, channelBufferSize)
    done := make(chan bool)

    // Define the callback chain
    callbackChain := NewCallbackChain(
        func(data Data) Data {
            // Simulate processing time
            fmt.Printf("Processing data: %d\n", data)
            return data + 10
        },
        func(data Data) Data {
            return data * 2
        },
    )

    // Number of producer and consumer goroutines
    numProducers := 1
    numConsumers := 1

    // Add workers to the wait group
    wg.Add(numProducers + numConsumers)

    // Start consumers
    for i := 0; i < numConsumers; i++ {
        go BackpressureConsumer(&wg, input, output, callbackChain)
    }

    // Start producers
    for i := 0; i < numProducers; i++ {
        go BackpressureProducer(&wg, input, done)
    }

    // Wait for producers to finish
    wg.Wait()

    // Close the output channel to signal consumers to finish
    close(done)

    // Wait for consumers to finish
    wg.Wait()

    // Close the output channel
    close(output)

    // Final results
    fmt.Println("Final results:")
    for result := range output {
        fmt.Println(result)
    }
}