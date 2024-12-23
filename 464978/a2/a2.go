package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Data structure
type Data struct {
	Value int
}

// Source stage
type Source struct {
	out  chan Data
	quit chan struct{}
	wg   *sync.WaitGroup
}

func NewSource(bufferSize int) *Source {
	return &Source{
		out:  make(chan Data, bufferSize),
		quit: make(chan struct{}),
		wg:   &sync.WaitGroup{},
	}
}

func (s *Source) Start() {
	s.wg.Add(1)
	go s.process()
}

func (s *Source) process() {
	defer s.wg.Done()
	rand.Seed(time.Now().UnixNano())
	for {
		select {
		case <-s.quit:
			close(s.out)
			return
		case s.out <- Data{Value: rand.Intn(100)}:
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond) // Simulate variable production time
		default:
			fmt.Println("Source: Buffer full, waiting.")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (s *Source) Close() {
	close(s.quit)
	s.wg.Wait()
}

// Processor stage
type Processor struct {
	in  chan Data
	out chan Data
	wg  *sync.WaitGroup
}

func NewProcessor(in chan Data, bufferSize int) *Processor {
	return &Processor{
		in:  in,
		out: make(chan Data, bufferSize),
		wg:  &sync.WaitGroup{},
	}
}

func (p *Processor) Start() {
	p.wg.Add(1)
	go p.process()
}

func (p *Processor) process() {
	defer p.wg.Done()
	for data := range p.in {
		processedData := Data{Value: data.Value * 2}
		select {
		case p.out <- processedData:
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond) // Simulate variable processing time
		default:
			fmt.Println("Processor: Buffer full, waiting.")
			time.Sleep(100 * time.Millisecond)
		}
	}
	close(p.out)
}

func (p *Processor) Close() {
	p.wg.Wait()
}

// Sink stage
type Sink struct {
	in chan Data
	wg *sync.WaitGroup
}

func NewSink(in chan Data) *Sink {
	return &Sink{
		in: in,
		wg: &sync.WaitGroup{},
	}
}

func (s *Sink) Start() {
	s.wg.Add(1)
	go s.process()
}

func (s *Sink) process() {
	defer s.wg.Done()
	for data := range s.in {
		fmt.Println("Sink: Received data", data.Value)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // Simulate variable processing time
	}
}

func (s *Sink) Close() {
	s.wg.Wait()
}

func main() {
	sourceBufferSize := 10
	processorBufferSize := 20

	source := NewSource(sourceBufferSize)
	processor := NewProcessor(source.out, processorBufferSize)
	sink := NewSink(processor.out)

	source.Start()
	processor.Start()
	sink.Start()

	time.Sleep(5 * time.Second) // Run for 5 seconds

	source.Close()
	processor.Close()
	sink.Close()

	fmt.Println("Pipeline completed.")
}