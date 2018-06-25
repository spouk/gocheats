package library

import (
	"math/rand"
	"log"
	"os"
	"time"
	"sync"
)

type ExampleSlice struct {
	Log       *log.Logger
	Stock     []int
	sync.WaitGroup
	chanEnd   chan bool
	chanStart chan bool
}

func NewExampleSlice() *ExampleSlice {
	return &ExampleSlice{
		Log:       log.New(os.Stdout, "[slice]", log.Lshortfile|log.Ltime),
		chanEnd:   make(chan bool),
		chanStart: make(chan bool),
	}
}
func (e *ExampleSlice) RandomElement(count int, period int) []int {
	defer func() {
		e.TimeTrackerWrapper("RandomElementRecursive")()
	}()
	var result []int
	for x := 0; x < count; x++ {
		result = append(result, rand.Intn(period))
	}
	return result
}
func (e *ExampleSlice) RandomElementRecursive(count, period int, result *[]int) {
	defer func() {
		e.TimeTrackerWrapper("RandomElementRecursive")()
	}()
	if count != 0 {
		*result = append(*result, rand.Intn(period))
		e.Log.Print(result)
		count --
		e.RandomElementRecursive(count, period, result)
	}
	return
}
func (e *ExampleSlice) TimeTracker(start time.Time, name string) {
	e.Log.Printf("Time execution `%s` function: %v\n", name, time.Since(start))
}
func (e *ExampleSlice) TimeTrackerWrapper(name string) func() {
	start := time.Now()
	return func() {
		e.Log.Printf("%s time %v\n", name, time.Since(start))
	}
}
func (e *ExampleSlice) RunWorker(count int) {
	for x := 0; x < count; x ++ {

	}
}
func (e *ExampleSlice) worker(id int) {
	e.Log.Printf("worker %d start", id)
	defer func() {
		e.Log.Printf("worker %d end", id)
	}()
	<- e.chanStart
	for {
		select {
		case <- e.chanEnd:
			return
		default:
			if len(e.Stock) == 0 {
				return
			} else {
				e.Stock = append(e.Stock[:0], e.Stock[:0]...)
			}
		}
	}
}
