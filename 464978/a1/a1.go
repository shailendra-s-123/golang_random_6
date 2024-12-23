package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Processing stage interface
type ProcessingStage interface {
	Start()
	Close()
}

// Sample data structure
type Data struct {
	Value int
}

// Source stage
type Source struct {
	out chan Data
	wg  *sync.WaitGroup
}

func NewSource(bufferSize int) *Source {
	return &Source{
		out: make(chan Data, bufferSize),
		wg:  &sync.WaitGroup{},
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
		case s.out <- Data{rand.Intn(100)}:
		default:
			fmt.Println("Source: Buffer full, waiting.")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (s *Source) Close() {
	close(s.out)
	s.wg.Wait()
}

// Processor stage
type Processor struct {
	in  chan Data
	out chan Data
	wg  *sync.WaitGroup
}

func NewProcessor(in chan Data, out chan Data, bufferSize int) *Processor {
	return &Processor{
		in:  in,
		out: out,
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
		processedData := Data{data.Value * 2}
		select {
		case p.out <- processedData:
		default:
			fmt.Println("Processor: Buffer full, waiting.")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (p *Processor) Close() {
	close(p.out)
	p.wg.Wait()
}

// Sink stage
type Sink struct {
	in  chan Data
	wg  *sync.WaitGroup
	done chan bool
}

func NewSink(in chan Data) *Sink {
	return &Sink{
		in:  in,
		wg:  &sync.WaitGroup{},
		done: make(chan bool),
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
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // Simulate processing time
	}
	s.done <- true
}

func (s *Sink) Close() {
	close(s.in)
	<-s.done
	s.wg.Wait()
}

func main() {
	sourceBufferSize := 10
	processorBufferSize := 20

	source := NewSource(sourceBufferSize)
	processor := NewProcessor(source.out, make(chan Data, processorBufferSize), processorBufferSize)
	sink := NewSink(processor.out)

	go source.Start()
	go processor.Start()
	go sink.Start()

	time.Sleep(5 * time.Second) // Run for 5 seconds

	source.Close()
	processor.Close()
	sink.Close()

	fmt.Println("Pipeline completed.")
}