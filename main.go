package main

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type RingBuffer struct {
	data       []*Data
	size       int
	nextRead   int
	lastInsert int
}

type Data struct {
	Stamp time.Time
	Value string
}

// constructor
func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		data:       make([]*Data, size),
		size:       size,
		nextRead:   0,
		lastInsert: -1,
	}
}

// Insert data into the RingBuffer
func (r *RingBuffer) Insert(input Data) {
	r.lastInsert = (r.lastInsert + 1) % r.size

	r.data[r.lastInsert] = &input

	if r.nextRead == r.lastInsert {
		r.nextRead = (r.nextRead + 1) % r.size
	}
}

// method to write to the content of the RingBuffer
func (r *RingBuffer) Emit() []*Data {
	output := []*Data{}

	for {
		if r.data[r.nextRead] != nil {
			output = append(output, r.data[r.nextRead])
			r.data[r.nextRead] = nil
		}

		if r.nextRead == r.lastInsert || r.lastInsert == -1 {
			break
		}

		r.nextRead = (r.nextRead + 1) % r.size
	}

	return output
}

// little test using the main func
func main() {
	rb := NewRingBuffer(5)
	currentRune := 'a' - 1

	fmt.Println("EMPTY TEST:")
	spew.Dump(rb.Emit())

	fmt.Println("FULL TEST:")
	for i := 0; i < 10; i++ {
		currentRune++
		rb.Insert(Data{
			Stamp: time.Now(),
			Value: string(currentRune),
		})
	}
	spew.Dump(rb.Emit())

	fmt.Println("PARTIAL TEST:")
	for i := 0; i < 8; i++ {
		currentRune++
		rb.Insert(Data{
			Stamp: time.Now(),
			Value: string(currentRune),
		})
	}
	spew.Dump(rb.Emit())
}
