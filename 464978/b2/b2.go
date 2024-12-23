package main  
import (  
    "fmt"
    "math/rand"
    "sync"
    "time"
)

const (
    minBufferSize  = 10
    maxBufferSize  = 100
    bufferSizeStep = 10
)

// Data structure
type Data struct {
    Value int
}

// Producer stage
type Producer struct {
    out chan Data
    wg  *sync.WaitGroup
}

func NewProducer() *Producer {
    return &Producer{
        out: make(chan Data, minBufferSize),
        wg:  &sync.WaitGroup{},
    }
}

func (p *Producer) Start() {
    p.wg.Add(1)
    go p.produce()
}

func (p *Producer) produce() {
    defer p.wg.Done()
    rand.Seed(time.Now().UnixNano())
    for {
        select {
        case <-time.After(time.Duration(rand.Intn(1000)) * time.Millisecond):
            data := Data{Value: rand.Intn(100)}
            p.out <- data
            fmt.Println("Producer: Sent data", data.Value)
        default:
            // If the buffer is full, adjust the buffer size dynamically
            p.adjustBufferSize(true)
            time.Sleep(100 * time.Millisecond)
        }
    }
}

func (p *Producer) adjustBufferSize(increase bool) {
    currentSize := cap(p.out)
    if increase && currentSize < maxBufferSize {
        newSize := currentSize + bufferSizeStep
        if newSize > maxBufferSize {
            newSize = maxBufferSize
        }
        newChan := make(chan Data, newSize)
        for data := range p.out {
            newChan <- data
        }
        p.out = newChan
        fmt.Println("Producer: Increased buffer size to", newSize)
    } else if !increase && currentSize > minBufferSize {
        newSize := currentSize - bufferSizeStep
        if newSize < minBufferSize {
            newSize = minBufferSize
        }
        newChan := make(chan Data, newSize)
        for data := range p.out {
            newChan <- data
        }
        p.out = newChan
        fmt.Println("Producer: Decreased buffer size to", newSize)
    }
}

func (p *Producer) Close() {
    close(p.out)
    p.wg.Wait()
}

// Consumer stage
type Consumer struct {
    in chan Data
    wg *sync.WaitGroup
}

func NewConsumer(in chan Data) *Consumer {
    return &Consumer{
        in: in,
        wg: &sync.WaitGroup{},
    }
}

func (c *Consumer) Start() {
    c.wg.Add(1)
    go c.consume()
}

func (c *Consumer) consume() {
    defer c.wg.Done()
    for data := range c.in {
        fmt.Println("Consumer: Received data", data.Value)
        time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // Simulate processing time
    }
}

func (c *Consumer) Close() {
    c.wg.Wait()
}

func main() {
    producer := NewProducer()
    consumer := NewConsumer(producer.out)

    producer.Start()
    consumer.Start()

    time.Sleep(10 * time.Second) // Run for 10 seconds

    producer.Close()
    consumer.Close()

    fmt.Println("Pipeline completed.")
} 