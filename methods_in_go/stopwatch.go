package main

import (
	"fmt"
	"time"
)

// Секундомер хранит начальное время time.Time, промежуточные отсечки []time.Duration
type Stopwatch struct {
	Start_time time.Time
	Results    []time.Duration
}

func (s *Stopwatch) Start() {
	s.Start_time = time.Now()
	s.Results = s.Results[:0]
}

func (s *Stopwatch) SaveSplit() {
	s.Results = append(s.Results, time.Since(s.Start_time))
}

func (s Stopwatch) GetResults() []time.Duration {
	return s.Results
}

func main() {
	sw := Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()
	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()
	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())

	sw.Start()
	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())
}
