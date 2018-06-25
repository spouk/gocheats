package main

import (
	"sync"
	"fmt"
	"time"
)

type CoreTeester struct {
	sync.WaitGroup
	command chan string
	timer   *time.Timer
}

func NewCoreTeester() *CoreTeester {
	return &CoreTeester{
		command: make(chan string, 3),
		timer:   time.NewTimer(time.Duration(time.Second * 30)),
	}
}
func (c *CoreTeester) worker(name interface{}) {
	fmt.Printf(fmt.Sprintf("worker %v\n", name))
	defer func() {
		c.Done()
		fmt.Printf(fmt.Sprintf("worker %v end\n", name))
	}()
	for {
		select {
		case command := <-c.command:
			if command == "exit" {
				return
			}
		default:
			time.Sleep(time.Duration(time.Second * 10))
			fmt.Printf(fmt.Sprintf("worker %v awaiking...\n", name))
		}
	}
}
func main() {
	c := NewCoreTeester()
	go func(){
		defer func(){
			c.Done()
		}()
		for {
			select {
			case <-c.timer.C:
				fmt.Printf("End timer ON\n")
				for i:=0; i<2; i++{
					c.command <- "exit"
				}
			}
		}
	}()
	for i := 0; i < 3; i++ {
		go c.worker(i)
		c.Add(1)
	}
	c.Wait()
	fmt.Printf("The end \n")
}
